# gstrings

gstrings is a more capable, UTF-8 aware version of the standard strings utility.
For full documentation, fork from [robpike/strings](https://github.com/robpike/strings).

```sh
$ gstrings -search ehprquli@iturq.info data/bingoo/t4.ibd
data/bingoo/t4.ibd:#105983:	皇甫艩嵯山东省潍坊市吢耽路7494号潯條小区15单元2282室ehprquli@iturq.info178195239864851373519850729873X
data/bingoo/t4.ibd:#106135:	郎熚黑海南省海口市肂憭路7342号磔奐小区17单元1329室ehprquli@iturq.info1769077249868331188197701111190
```

```shell
$ gstrings t3.ibd | tail
t3.ibd:#325119:	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
t3.ibd:#325375:	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
t3.ibd:#325631:	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
t3.ibd:#325887:	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
t3.ibd:#326143:	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
t3.ibd:#326399:	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
t3.ibd:#326655:	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
t3.ibd:#326911:	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
t3.ibd:#327167:	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
t3.ibd:#327423:	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa

$ gstrings t4.ibd | tail
t4.ibd:#105101:	莘盗牁甘肃省定西市撅齗路6272号箲蛀小区4单元1207室aacrqnhl@ygjtl.me180880142164351575820160114069X
t4.ibd:#105247:	成垄纆云南省丽江市皘疬路416号苁节小区1单元1376室mjilglyl@ibhad.vip1533230689186340644198005194531
t4.ibd:#105393:	扶鮯銦陕西省咸阳市滍稈路6543号絗欀小区8单元1925室hdxwftmr@exzal.com1572786679698444319198206210387
t4.ibd:#105540:	曲陞曞安徽省铜陵市嗌鉆路2477号豦袌小区12单元1839室rbidijxa@phiuq.co1703951692479311443199307013911
t4.ibd:#105687:	广亍円江西省吉安市钬麛路637号簈獕小区4单元1478室xcimqzwc@yccum.store180854600904734390919770811023X
t4.ibd:#105835:	师娘繳江苏省扬州市輨迃路1116号噜鶖小区18单元675室skahwesz@nnfad.info1528457027138139643197107164331
t4.ibd:#105983:	皇甫艩嵯山东省潍坊市吢耽路7494号潯條小区15单元2282室ehprquli@iturq.info178195239864851373519850729873X
t4.ibd:#106135:	郎熚黑海南省海口市肂憭路7342号磔奐小区17单元1329室ehprquli@iturq.info1769077249868331188197701111190
t4.ibd:#106284:	通蒫冒广东省汕尾市啞鵟路3432号櫳峣小区10单元557室ehprquli@iturq.info1808537856755372400199001052477
t4.ibd:#106432:	印迎値云南省昭通市蠬咏路2554号廥溪小区18单元1055室ehprquli@iturq.info1503766057953544390200511305364
```

## find by hex raw bytes

```shell
$ gstrings --offset --hex 2f6c69 ~/RustroverProjects/hash_macro/target/release/hash_macro                                                          
strings: Found at 1667, Offset 1665: 73722f6c6962
```
