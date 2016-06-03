f = open("Hashmat_the_Brave_Warrior_input.md",'r')
o = open("Hashmat_the_Brave_Warrior_output.md",'w')

def Hashmat_the_Brave_Warrior(warriors,opponents):

    if isinstance( warriors, int ) and isinstance( opponents, int ):
        return opponents-warriors
    return False

while True:
    line = f.readline()
    if not line:
        break
    elems = line.partition(" ")
    print(type(elems[0]),elems[0],elems[2],elems)
    diff = Hashmat_the_Brave_Warrior(int(elems[0]),int(elems[2]))
    o.write(str(diff)+"\n")


f.close()
o.close()
