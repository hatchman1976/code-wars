import Maskify

def test_Maskify():
    cc = '123'
    r = Maskify.maskify(cc)
    #test.describe("masking: {0}".format(cc))
    #test.it("{0}  matches  {1}".format(cc, r))
    assert (r, cc)

def test_MaskBigger():
    cc = 'SF$SDfgsd2eA'
    r = Maskify.maskify(cc)

    assert (r, '########d2eA')