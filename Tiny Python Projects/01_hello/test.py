#!/usr/bin/env python3
# unfortunately, had to modify testcases to work with windows cmd :(
"""tests for hello.py"""

import os
from subprocess import getstatusoutput, getoutput

prg = "./hello.py"
executable = r"C:\Program Files\Git\git-bash.exe"

# --------------------------------------------------
def test_exists():
    """exists"""

    assert os.path.isfile(prg)


# --------------------------------------------------
def test_runnable():
    """Runs using python3"""

    out = getoutput(f"python {prg}")
    assert out.strip() == "Hello, World!"


# --------------------------------------------------
def test_executable():
    """Says 'Hello, World!' by default"""

    out = getoutput(f"python {prg}")
    assert out.strip() == "Hello, World!"


# --------------------------------------------------
def test_usage():
    """usage"""

    for flag in ["-h", "--help"]:
        rv, out = getstatusoutput(f"python {prg} {flag}")
        assert rv == 0
        assert out.lower().startswith("usage")


# --------------------------------------------------
def test_input():
    """test for input"""

    for val in ["Universe", "Multiverse"]:
        for option in ["-n", "--name"]:
            rv, out = getstatusoutput(f"python {prg} {option} {val}")
            assert rv == 0
            assert out.strip() == f"Hello, {val}!"
