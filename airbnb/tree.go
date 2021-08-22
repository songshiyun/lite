package main

/**
sample input:
/bin/usr/bash
/var/test/.ssh
/var/log/wifi.log
/opt
/xyz

sample output:
/
+-- bin
|   +-- usr
|       +-- bash
+-- opt
+-- var
|   +-- log
|   |   +-- wifi.log
|   +-- test
|       +-- .ssh
+-- xyz
*/
func Tree(string2 string) string {

	return string2
}
