import Bitcounting

def test_bitcount():
    assert (Bitcounting.countBits(0), 0)
    assert (Bitcounting.countBits(4), 1)
    assert (Bitcounting.countBits(7), 3)
    assert (Bitcounting.countBits(9), 2)
    assert Bitcounting.countBits(10) == 2