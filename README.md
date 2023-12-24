# How to use
You need to have python installed (I used python 3.8), as well as golang (I used go 1.21).

You can either copy-paste the input for the day XX in ./dayXX/input.txt (for example for the day 4 it would be in ./day04/input.txt), or you can just copy paste your session token in ./session.txt and the script will get it for you.
You can also copy paste the example values for the day XX in ./dayXX/example.txt to make sure it matches the given solution.

Then just run `python run.py` to get the solution for the last day available.

You can add the flag `--dry-run` to run only your functions on the data in example.txt (doing so won't even try to read or download input.txt).

You can add the flag `-day=X` to specify a day. For example in you are currently in the folder ./day08/ and you want to run the your functions of the day 3 but only with the example, run `python ../run.py -day=3 --dry-run`

You can add the flag `--no-part-1` to skip the part 1 and only run the part 2 (usefull when the part 1 takes too long).

