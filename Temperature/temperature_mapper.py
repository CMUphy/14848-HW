import sys

def checkQuality(flag):
    if flag == 0 or flag == 1 or flag == 4 or flag == 5 or flag == 9 :
        return True
    else: return False

for line in sys.stdin:
    line = line.strip()
    temperature = int(line[87:92])
    quality = int(line[92])
    if ((temperature != 9999) and checkQuality(quality)):
        print('%s\t%d' % (line[15:23], int(line[87:92])))
