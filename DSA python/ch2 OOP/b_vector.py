class Vector:
    """Represent a vector in a multidimensional space"""

    def __init__(self, d):
        """Create d-dimensional vector of zeros."""
        self._coords = [0] * d

    def __len__(self):
        """Return the dimension of the vector."""
        return len(self._coords)

    def __getitem__(self, j):
        """Return jth coordinate of vector"""
        return self._coords[j]

    def __setitem__(self, j, val):
        """Set jth coordinate of vector to given value"""
        self._coords[j] = val

    def __add__(self, other):
        """Return sum of two vectors"""
        if len(self) != len(other):  # relies on len method
            raise ValueError("dimensions must be same")
        result = Vector(len(self))  # start with vector of zeros
        for j in range(len(self)):
            result[j] = self[j] + other[j]
        return result

    def __eq__(self, other: object) -> bool:
        """Return True if vector has same coordinates as other."""
        return self._coords == other._coords

    def __ne__(self, other: object) -> bool:
        """Return True if vector differs from other."""
        return not self == other  # rely on existing eq definition

    def __str__(self) -> str:
        """Produce string representation of vector."""
        return f"<{str(self._coords)[1:-1]}>"  # adapt list representation


if __name__ == "__main__":
    v = Vector(5)  # construct five-dimensional <0, 0, 0, 0, 0>
    v[1] = 23  # <0, 23, 0, 0, 0> (based on use of setitem )
    v[-1] = 45  # <0, 23, 0, 0, 45> (also via setitem )
    print(f"v: {v}")
    print(f"v4: {v[4]}")  # print 45 (via getitem )
    u = v + v  # <0, 46, 0, 0, 90> (via add )
    print(f"u: {u}")  # print <0, 46, 0, 0, 90>
    total = 0
    for entry in v:  # implicit iteration via len and getitem
        total += entry

