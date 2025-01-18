from typing import List
from enum import Enum
import logging

logger  = logging.getLogger(__name__)

class ClassRegistry(str, Enum):
    QUERY_OPTIMIZER = "query_optimizer"
    LANGUAGE_MODEL = "language_model"

__registry = {}

def get_registered(category: ClassRegistry) -> List[str]:
    global __registry 
    return __registry.get(category, [])

def register(category: ClassRegistry, cls):
    global __registry

    if category not in __registry:
        __registry[category] = []
        
    if not hasattr(cls, '__name__'):
        logger.error(f"Class {cls} does not have __name__ attribute")
        return False

    logger.info(f"Registering {cls.__name__} as a {category}")
    __registry[category].append(cls)
    return True

def register_decorator(category: ClassRegistry):
    def decorator(cls):
        register(category, cls)
        return cls
    return decorator
    
def get_cls(category: ClassRegistry, name: str):
    global __registry

    for cls in __registry.get(category, []):
        if cls.__name__ == name:
            return cls

    return None