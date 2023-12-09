import os
import pathlib
import subprocess
import sys

file_path = pathlib.Path(__file__).parent.resolve()
os.chdir(file_path)
files = [ f.path for f in os.scandir("./")]

if (not "./run.py" in files):
    print(files)
    print("Error with file location ", file_path)
    exit()
   
subfolders = [ f.path[2:] for f in os.scandir("./") if f.is_dir() and len(f.path)>2 ]
days = [ name for name in subfolders if name.startswith("day") ]

flags = []
dayFlag = 0
for arg in sys.argv[1:]:
    if arg.startswith("-day="):
        try:
            day = int(arg[5:])
            if day > 0 and day <= 25:
                dayFlag = day
        except Error as err:
            print(err)
    elif arg == "--dry-run":
        flags.append(arg)
print(flags)

lastDay = 0
dayName = ""
for day in days:
    try:
        dayNum = int(day[3:])
        if dayNum > lastDay and dayNum <= 25:
            dayName = day
            lastDay = dayNum
        if dayNum == dayFlag:
            dayName = day
            lastDay = dayNum
            break
    except Error as err:
        print(err)
if dayFlag > 0 and dayFlag != dayNum:
    print("day "+str(dayFlag)+" not found")
    exit()

if not lastDay > 0:
    print("No day found")
    exit()

flags.append("-day="+str(lastDay))

os.chdir(dayName)
pluguinBuildResult = subprocess.run(["go build -buildmode=plugin " +dayName+".go"],shell=True,capture_output = True, text=True)
print(pluguinBuildResult.stdout)
if len(pluguinBuildResult.stderr) > 0:
    print("error")
    print(pluguinBuildResult.stderr)
    exit()
os.chdir("..")
cmdToRun = "go run main.go "+str.join(" ",flags)+" "+dayName+"/"+dayName+".so"
print("> "+cmdToRun+"\n")
subprocess.run([cmdToRun],shell=True)


