import argparse
import os


def get_args():
    """Get command-line arguments"""

    parser = argparse.ArgumentParser(
        description="Shout out", formatter_class=argparse.ArgumentDefaultsHelpFormatter
    )

    parser.add_argument("text", metavar="text", type=str, help="Input string or file")

    parser.add_argument(
        "-o", "--outfile", help="Output filename", metavar="str", type=str, default=""
    )

    # check if the text supplied is an actual file
    args = parser.parse_args()

    if os.path.isfile(args.text):
        args.text = open(args.text).read().strip()

    return args


def main():
    args = get_args()

    if args.outfile:
        with open(args.outfile, "wt") as f:
            f.write(args.text.upper())
    else:
        print(args.text.upper())


if __name__ == "__main__":
    main()
