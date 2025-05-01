// SPDX-License-Identifier: Apache-2.0

// Package internal provides shared infrastructure for Otto.
package internal

import (
	"fmt"
)

// OttoASCII is the ASCII art representation of Otto the otter
// Generated from bin/otto.png using jp2a with width=80 --border option
const OttoASCII = `+--------------------------------------------------------------------------------+
|                                                                                |
|                                   ;ol:'.                                       |
|                                ...;cxxxxlclllcc:;'..                           |
|                      .,,'. .,cdxl,coxxxxxxxxxxxxxxxxoc,.  .';:;,.              |
|                    .lcol..cxxxxxxxxxxxxxxxxxxxxxxxxxxxxx:cxxxdood'             |
|                    :dc .lxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxd,.,:do             |
|                    .c.;dxd:;:dxxxxxxxxxxxxxxd:;;:dxxxxxxxxxl';oxx;             |
|                      cxxdlododxxxxxxxxxxxxxxdoooloxxxxxxxxxxdlxl.              |
|                     .'''''''.,lxxxxxxxxxxo;.'',,''.':dxxxxxxxxd;               |
|                    .;oxxxxxxd:.'dxxxxxxo'.cdxxxxxxxl,.;xxxxxxxxx.              |
|                    kxxxd;..';dd.'xxxdxo.,dxx:,...,dxxo.,xxxxxxxxl              |
|                   :kxxd.  ,Mk'x: ,...'..dxx;cMd   .xxxc ..;xxxxxx.             |
|                   lxxxx    '..x: coddl..xxx. '     dxxl cdxxxxxxx;             |
|                   .Okkx,     lo.,:;;:ld lxx:      .xxx,.xkkxxxxxxl             |
|       ..         dl.xXNKl;,:l;.d,    '0x.cOOl'..'cOKO:'KNNNXOxoxxd             |
|    .ldlxl.oc.   'NKOc,;::,,,;dXNNKxdKNNNXd,,:llloo:,,o0KXNNKkxclxo             |
|   .x,...o,o00..'.,;cok0XXXNNk:k0Od;,cxO00x;k0kxdxkk000000XNOKNOxx;             |
|   x.,:;.,l:00;'dddol;'',:lk0NXkx..,,.''''.xNNNNNNNXKKKKKXNNKxo:do              |
|  'o.::;.;:l00;.cccclooddol;.,KNNk..:dkkkd'XNNNNNNNNNNNNNNNNNXkdl.              |
|  ::'::,.x.k00.,ccccccccccllo..',c;.,col;;KNNNNNNNNNNNNNNNNNKxo,                |
|  'o.;,.o;l00c.cccccccccccccc.;ool:;'..'lOX0KNNNNNNNNNNNNX0d:.                  |
|   od:ld,lOO: ;:::cccccccccc: :cccccclol:;,. '::KNNNXK0dc,.                     |
|      ..coc. ..'',,,'',;;:::.':::ccccccccclo.cK:'kxxxxoloc              ;c.     |
|                  ..c,.c... ...'',,,;;;;,'''.:O;'xo::xxxxx'             oxd'    |
|                   oxoxx;o ::..:,,'.....,cddoc;',,;cdxxxxxd.            dxxd.   |
|                   :xxxxxx:ll.OK00OOOx:.loxxxxdxxxxxxxxxxxxc           .xxxx;   |
|                    ,oxxxxdl.:NNNNNNXXK,.cxddxxxxxxxxxxxdxxx'          :xxxxl   |
|                       ,::;' KNNNNNNNNNNx;,:odxxxxxxxdc'cxxxo         .dxxxxo   |
|                            :NNNNNNNNNNNNNKxc:;,,,,,,':dxxxxx;       .dxxxxxo   |
|                            0NNNNNNNNNNNNNNNNNNNOc;:oxxxxxxxxd.     ;dxxxxxx:   |
|                         ',.NNNNNNNNNNNNNNNNNNXl';oxxxxxxxxxxx,  .,oxxxxxxxd.   |
|                       .lx:'NNNNNNNNNNNNNNNNNk.cxxxxxxxxxxxxxxc.ldxxxxxxxxd;    |
|                       cxx:'NNNNNNNNNNNNNNNN0.oxxxxxxxxxxxxxxxl.xdxxxxxxxd,     |
|                       dxxd.KNNNNNNNNNNNNNNN,;dxxxxxxxxxxxxxxxc.dxxxxddoc.      |
|                       :xxx;'XNNNNNNNNNNNNNN.cxxxxxxxxxxxxxxxx',ddddolc'        |
|                        cxdd,'0NNNNNNNNNNNNN:'dxxxxxxxxxxxxxxc.clccc;.          |
|                         .;cdc.,xXNNNNNNNNNNX.;ddxxxxdxxxxxd:.:c:,.             |
|                      .':oolodoc,.,:oxkO0000Ox..::ccodxxxxo.                    |
|                      ''d:,odxdo;'           .::dddldxxxxd,                     |
|                                             . col.clc:,                        |
|                                                                                |
+--------------------------------------------------------------------------------+
`

// DisplayBanner prints the Otto ASCII art banner to the console
func DisplayBanner() {
	fmt.Print(OttoASCII)
}

// DisplayBannerWithVersion prints the Otto ASCII art banner with version information
func DisplayBannerWithVersion(version string) {
	fmt.Print(OttoASCII)
	fmt.Printf("Version: %s\n\n", version)
}
