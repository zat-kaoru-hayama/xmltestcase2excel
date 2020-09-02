xmltestcase2excel
=================

```
<xml>
<item>
    <h>TEST CASE (1) TITLE</h>
    <ope>User operations</ope>
    <status>OK or NG</status>
</item>
<item>
    <h>TEST CASE (2) TITLE</h>
    <ope>User operations</ope>
    <status>OK or NG</status>
</item>
</xml>
```

```
xmltestcase2excel [-template TEMPLATE.XLSX] [-output OUTPUT.XLSX] source.xml
```

* `item>h` to B3,B4,B5...
* `item>ope` to C3,C4,C5...
* `item>status` to E3,E4,E5...
