import collections

def is_isogram(string):
   return len([i < 1 for i in collections.Counter(string.lower()).values()]) == len(string)