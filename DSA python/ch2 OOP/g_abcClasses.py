# An abstract base class akin to collections.Sequence.

# This implementation relies on two advanced Python techniques. The first is that
# we declare the ABCMeta class of the abc module as a metaclass of our Sequence
# class. A metaclass is different from a superclass, in that it provides a template for
# the class definition itself. Specifically, the ABCMeta declaration assures that the
# constructor for the class raises an error.

# The second advanced technique is the use of the @abstractmethod decorator
# immediately before the len and getitem methods are declared. That
# declares these two particular methods to be abstract, meaning that we do not provide
# an implementation within our Sequence base class, but that we expect any concrete
# subclasses to support those two methods. Python enforces this expectation,
# by disallowing instantiation for any subclass that does not override the abstract methods
# with concrete implementations.


from abc import ABCMeta, abstractmethod  # need these definitions


class Sequence(metaclass=ABCMeta):
    """Our own version of collections.Sequence abstract base class."""

    @abstractmethod
    def len(self):
        """Return the length of the sequence."""

    @abstractmethod
    def getitem(self, j):
        """Return the element at index j of the sequence."""

    def contains(self, val):
        """Return True if val found in the sequence; False otherwise."""
        for j in range(len(self)):
            if self[j] == val:  # found match
                return True
        return False

    def index(self, val):
        """Return leftmost index at which val is found (or raise ValueError)."""
        for j in range(len(self)):
            if self[j] == val:  # leftmost match
                return j
        raise ValueError("value not in sequence")  # never found a match

    def count(self, val):
        """Return the number of elements equal to given value."""
        k = 0
        for j in range(len(self)):
            if self[j] == val:  # found a match
                k += 1
        return k
