import os
import glob
import subprocess


def dir_walk(path, name):
    if not glob.glob(path + "/*go.mod*"):
        return

    #print('+' + path)
    print('[' + name + ']')
    cmd = 'cd ' + path + ';' + 'go test -v;' 
    #cmd = 'cd ' + path + '; pwd;'  
    #print(cmd)
    subprocess.run(cmd , shell=True)


for curDir, dirs, files in os.walk("."):
    if '.git' in dirs:
        dirs.remove('.git')
    for dir in dirs:
        path = os.path.join(curDir, dir)
        #print(path)
        dir_walk(path, dir)
