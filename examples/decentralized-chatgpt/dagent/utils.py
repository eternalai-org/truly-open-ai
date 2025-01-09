import datetime 
import os
from typing import Optional
import logging
from .models import InferenceResult
import queue
from singleton_decorator import singleton

logger = logging.getLogger(__name__)

def formated_utc_time():
    return datetime.datetime.utcnow().strftime("%Y-%m-%dT%H:%M:%S.%fZ")

def get_script_dir(ee = __file__):
    return os.path.dirname(os.path.realpath(ee))

@singleton
class SimpleCacheMechanism(object):
    MAX_CACHE_ITEMS = 2048

    def __init__(self, *args, **kwargs):
        self._log = {}
        self._que = queue.Queue()

    def commit(self, result: InferenceResult) -> InferenceResult:
        self._log[result.id] = result
        self._que.put(result.id)

        while len(self._log) > self.MAX_CACHE_ITEMS:
            top = self._que.get()
            self._log.pop(top)

        return result

    def get(self, id: str, default=None) -> Optional[InferenceResult]:
        return self._log.get(id, default)

from enum import Enum
import sys

class ConsoleColor(str, Enum):
    RED = '\033[1;31m'
    GREEN = '\033[1;32m'
    YELLOW = '\033[1;33m'
    BLUE = '\033[1;34m'
    COLOR_OFF = '\033[1;0m'
    WHITE = '\033[1;37m'
    JUST_BOLD = '\033[1m'

def print_error(msg):
    if type(msg) == bytes:
        msg = msg.decode('utf-8', errors = 'ignore')

    print_color(f"[ERROR] {str(msg)}".rstrip('\n \t'), ConsoleColor.RED, file = sys.stderr)

def print_warn(msg):
    if type(msg) == bytes:
        msg = msg.decode('utf-8', errors = 'ignore')

    print_color(f"[WARNING] {str(msg)}".rstrip('\n \t'), ConsoleColor.YELLOW, file = sys.stdout)

def print_color(msg, color: ConsoleColor = ConsoleColor.WHITE, file = None, end = '\n', flush=False):
    if file is None or not hasattr(file, 'write'):
        file = sys.stdout

    print(f'{color.value}{msg}{ConsoleColor.COLOR_OFF.value}', file=file, end=end, flush=flush)