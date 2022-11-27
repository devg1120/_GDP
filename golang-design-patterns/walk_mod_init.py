import os
import glob
import subprocess

def mod_init(path, name):
    print('       mod init ' + name)
    cmd = 'cd ' + path + '; pwd;' + 'go mod init ' + name +';'
    subprocess.run(cmd , shell=True)

def dir_walk(path, name):
    if not glob.glob(path + "/*.go*"):
        return

    print('+' + path)
    print('    ' + name)
    if glob.glob(path + "/go.mod*"):
        print('                              go.mod exist')
        return
    mod_init(path, name)


for curDir, dirs, files in os.walk("."):
    if '.git' in dirs:
        dirs.remove('.git')
    for dir in dirs:
        path = os.path.join(curDir, dir)
        #print(path)
        dir_walk(path, dir)
