#!/usr/bin/env python3
import argparse

import mido


def _run(args):
    inputs =  {i.lower().split()[0]: i for i in mido.get_input_names()}
    if args.list:
        for i in inputs.keys():
            print(i)
        return
    elif args.stream_from:
        try:
            with mido.open_input(inputs[args.stream_from]) as port:
                for msg in port:
                    print(msg)
        except KeyError:
            print(f'unavailable input "{args.stream_from}".')
            if inputs:
                print(f'available MIDI inputs are {", ".join(list(inputs.keys()))}.')
            else:
                print('there are no available MIDI devices!')


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--list", "-l", action="store_true", help="list existing MIDI devices.")
    parser.add_argument("--stream-from", "-s", help="stream MIDI messages from device.")
    
    args = parser.parse_args()

    _run(args)
