import os

def persistence(n):    
    count = 0
    
    while len(str(n)) > 1:
        nArray = list(str(n))
        n = 1
        for i in nArray:
            n *= int(i)
        count += 1
    
    return count
print(os.getcwd())

for dir_path, dir_names, file_names in os.walk(os.getcwd()):
    for f in file_names:
        print(f)
print(persistence(39))