# Golang net/rpc package bug?

This demo project seem shows a bug(or feature?) of net/rpc package.

## Detail Explaination

It seems net/rpc module cannot serilize 0 in the  reply struct.

In the demo code, if 0 can be serilized successfully, client should get reply.ID == 0, however, the result is 1.
