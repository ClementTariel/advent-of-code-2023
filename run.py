import os
import pathlib
import subprocess

file_path = pathlib.Path(__file__).parent.resolve()
os.chdir(file_path)
files = [ f.path for f in os.scandir("./")]

if (not "./run.py" in files):
    print(files)
    print("Error with file location ", file_path)
    exit()
   
subfolders = [ f.path[2:] for f in os.scandir("./") if f.is_dir() and len(f.path)>2 ]
days = [ name for name in subfolders if name.startswith("day") ]

lastDay = 0
dayName = ""
for day in days:
    if int(day[3:]) > lastDay and len(day)<6:
        dayName = day
        lastDay = int(day[3:])

if not lastDay > 0:
    print("No day found")
    exit()

os.chdir(dayName)
print(subprocess.run(["go build -buildmode=plugin " +dayName+".go"],shell=True))
os.chdir("..")
subprocess.run(["go run main.go "+dayName+"/"+dayName+".so"],shell=True)


