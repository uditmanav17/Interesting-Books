#!C:/Users/uditm/Anaconda3/python

# Purpose: Say hello

import argparse


def main():
    parser = argparse.ArgumentParser(description="Say hello")
    # compulsory params
    # parser.add_argument("name", help="Name to greet, compulsory")
    # optional params
    parser.add_argument(
        "-n", "--name", metavar="name", default="World", help="Name Optional"
    )
    parser.add_argument("-a", "--age", metavar="age", default=0, help="Age to print")
    args = parser.parse_args()

    if args.age:
        print(f"Hello, {args.name}! You are {args.age} years old.")
    else:
        print(f"Hello, {args.name}!")


if __name__ == "__main__":
    main()
