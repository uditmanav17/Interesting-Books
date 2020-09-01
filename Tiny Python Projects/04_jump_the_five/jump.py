#!/usr/bin/env python3
"""Jump the Five"""

import argparse


# --------------------------------------------------
def get_args():
    """Get command-line arguments"""

    parser = argparse.ArgumentParser(
        description="Jump the Five",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )

    parser.add_argument("text", metavar="str", help="Input text")

    return parser.parse_args()


# --------------------------------------------------
def main():
    """Make a jazz noise here"""

    args = get_args()
    mapper = dict(zip(map(str, range(10)), list("5987604321")))

    out_str = ""
    for char in args.text:
        out_str += mapper.get(char, char)
    print(out_str)
    # print(args.text.translate(str.maketrans(jumper)))


# --------------------------------------------------
if __name__ == "__main__":
    main()
