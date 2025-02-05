from typing import Dict
from queue import Queue


class TimeEstimation(Queue):
    DEFAULT_MAXSIZE = 10
    DEFAULT_ESTIMATION = 60

    def __init__(self, maxsize=0, default_estimation=60):
        super().__init__(maxsize)

        self.total = 0
        self.default_estimation = default_estimation

    def put(self, item, block=True, timeout=None):
        assert isinstance(item, (int, float)), "Item must be a number"

        prev = 0
        if self.qsize() == self.maxsize:
            prev = self.get()

        self.total += item - prev
        super().put(item, block, timeout)

    def estimate(self, block=True, timeout=None):
        if self.qsize() == 0:
            return self.default_estimation

        return self.total / self.qsize()


class ModelInferTimeEstimation:

    def __init__(self):
        self.models: Dict[str, TimeEstimation] = {}

    def prepare(self, model_id: str, maxsize=1000, default_estimation=60):
        if model_id not in self.models:
            self.models[model_id] = TimeEstimation(maxsize, default_estimation)

    def estimate(self, model_id: str):
        self.prepare(model_id)
        return self.models[model_id].estimate()

    def update(self, model_id: str, actual_time: float):
        self.prepare(model_id)
        self.models[model_id].put(actual_time)
