#!/usr/bin/env python
# -*- coding: utf-8 -*-
import sys

def main():
    ct = 0
    while ct < 1000000:
        print("hello")
        sys.stdout.flush()
        ct += 1


if __name__ == '__main__':
    main()