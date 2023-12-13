local dic = {}
math.randomseed(os.time())

i = 1
for line in io.lines("dic-zh.txt") do
    dic[i] = line
    i = i + 1
end
io.close()


eventSleep(1000)
for _ = 0, 10, 1 do
    click(269, 149)
    keyTaps({ "ctrl", "a" })
    keyTap("del")
    input(dic[math.random(i)])
    keyTap("enter")
    sleep(5)
end