Open Container Specification 
=======
Assume the upstream changes.

History of the Open assume upstream changes Container Initiative

Meaning of the Open Containeraaa Specification 
>>>>>>> ce5f2c9a1210757a94661ae02c8d98634670d31b

Introduction of the Open Container Testing framework

Plan of the Open Container Testing community





TODO-List (by seqence)

1. testcase server
1.1 parse & validation
    find the mandatory area
1.2 send request to test-server and container-server
1.3 add config file to tell the test-server daemon and container pool daemon address

2. test-server-daemon
2.1 monitor the request from testcase server
2.2 deliver to the underling HA/openStack...
2.3 return resource id/id-list to testcase server

3. container pool daemon
3.1 monitor the request from testcase server
3.2 deliver to the undeerling container hub
3.3 return resource id/id-list to testcase server

4. testcase server send new testcase order (with both id/id-list) to test-server-daemon
    to deploy the test

# go-hack
start to use golang


ssh-keygen -t rsa -f my-ssh-keys -C "google-ssh {"userName":"my-username@localhost","expireOn":"2018-12-04T20:12:00+0000"}"
