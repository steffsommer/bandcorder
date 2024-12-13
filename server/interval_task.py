import asyncio
from typing import Callable


class IntervalTask:

    def __init__(self, cycle_seconds: int, callback: Callable[[None], None]):
        self._interval_seconds = cycle_seconds
        self._callback = callback

    def start(self):
        loop = asyncio.get_event_loop()
        self._task = loop.create_task(self.execute_periodically())

    async def execute_periodically(self) -> None:
        while True:
            self._callback()
            await asyncio.sleep(self._interval_seconds)

    def stop(self):
        if not self._task:
            return
        self._task.cancel()
