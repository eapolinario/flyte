
�
&""core.type_system.custom_objects.wf �



x
x

y
y�

�
o1�
�	2�
�
definitions�*�
�
FlyteschemaSchema*}
M

properties?*=
;
remote_path,**

typestring

titleremote_path

typeobject

additionalProperties  
�
FlytefileSchemaq*o

typeobject
?

properties1*/
-
path%*#

titlepath

typestring

additionalProperties  
�
ResultSchema�*�

typeobject

additionalProperties  
�

properties�*�
[
schemaQ*O


field_many  

typeobject
)
$ref!#/definitions/FlyteschemaSchema
W
fileO*M


field_many  
'
$ref#/definitions/FlytefileSchema

typeobject
a
	directoryT*R
,
$ref$"#/definitions/FlytedirectorySchema


field_many  

typeobject
�
FlytedirectorySchemaq*o
?

properties1*/
-
path%*#

typestring

titlepath

additionalProperties  

typeobject
4
$schema)'http://json-schema.org/draft-07/schema#
$
$ref#/definitions/ResultSchemao1
�
o0�
�	2�
4
$schema)'http://json-schema.org/draft-07/schema#
#
$ref#/definitions/DatumSchema
�
definitions�*�
�
DatumSchema�*�

typeobject

additionalProperties  
�

properties�*�
'
y"* 

titley

typestring
(
x#*!

titlex

type	integer
c
z^*\

titlez

typeobject
:
additionalProperties"* 

typestring

titlezo0"N
n0
upload_result* 25
1"-core.type_system.custom_objects.upload_result "i
n1
download_result* 
res

n0o0"n027
3"/core.type_system.custom_objects.download_result "R
n2
	stringify* 

xx21
-")core.type_system.custom_objects.stringify "R
n3
	stringify* 

xy21
-")core.type_system.custom_objects.stringify "d
n4
add* 
x

n2o0
y

n3o0"n2"n32+
'"#core.type_system.custom_objects.add *
o0

n4o0*
o1

n0o0: 