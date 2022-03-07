from multiprocessing.sharedctypes import Value


var=str("lol")
integer=10 % 10
if var == "lol" :
    print("hi")
else :
    print("hi")
if 20 % 2 == 0:
    print("fuck")
elif 10 * 10 == 2 :
    print("non")

fruits = ["item1", "item2"]

fruits.append("item3")
print(fruits[2])

fruitsof = [
    [0 , 1 , 2] ,
    [0 , 1 , 2] 
]

dictionary = {"key":"Value" , "key2":"aner"}

print(fruitsof[1][1])

for item in fruits:
    print("hi")

for i in range(5):
    print("psp")
    

b=1
a=0
while 5 > a :
    print("hi")
    a=a+1
    
def function(x) :
    print(x) 
    global b
    b = 2
    return x 

function("cap")

print(dictionary["key2"])

dictitude2 = {"key":"value"}
dictitude3 = {"key":"value"}
dictitude = [dictitude2 , dictitude3 ]

dictitude5 = ["man" , "dick"]
dictitude6 = ["man" , "dicker "]
dictitude4 = {"man":dictitude5 , "female":dictitude6}

print(dictitude[0])
print(dictitude4["man"])

print(b)