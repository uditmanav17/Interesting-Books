# Python also helps by providing an automatic iterator implementation for any
# class that defines both len and getitem . To provide an instructive example
# of a low-level iterator, this Code demonstrates just such an iterator
# class that works on any collection that supports both len and getitem .
# This class can be instantiated as SequenceIterator(data). It operates by keeping an
# internal reference to the data sequence, as well as a current index into the sequence.
# Each time next is called, the index is incremented, until reaching the end of
# the sequence.


class SequenceIterator:
    """An iterator for any of Python s sequence types."""

    def __init__(self, sequence):
        """Create an iterator for the given sequence."""
        self.seq = sequence  # keep a reference to the underlying data
        self.k = -1  # will increment to 0 on first call to next

    def __next__(self):
        """Return the next element, or else raise StopIteration error."""
        self.k += 1  # advance to next index
        if self.k < len(self.seq):
            return self.seq[self.k]  # return the data element
        else:
            raise StopIteration()  # there are no more elements

    def __iter__(self):
        """By convention, an iterator must return itself as an iterator."""
        return self
