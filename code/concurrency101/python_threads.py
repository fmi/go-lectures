#!/usr/bin/env python

import time
import threading

def ticker():
    while 42:
        print("Tick!")
        time.sleep(1)

thread = threading.Thread(target=ticker)
thread.daemon = True
thread.start()

# Or using sublcassing:

class TickerThread(threading.Thread):
    def run(self):
        while 42:
            print("Tick!")
            time.sleep(1)

thread = TickerThread()
thread.start()
# ...
thread.join()
