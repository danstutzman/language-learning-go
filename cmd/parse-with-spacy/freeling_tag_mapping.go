/* Adapted from https://github.com/linhd-postdata/spacy-affixes/blob/develop/src/spacy_affixes/eagles.py

                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

   TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

   1. Definitions.

      "License" shall mean the terms and conditions for use, reproduction,
      and distribution as defined by Sections 1 through 9 of this document.

      "Licensor" shall mean the copyright owner or entity authorized by
      the copyright owner that is granting the License.

      "Legal Entity" shall mean the union of the acting entity and all
      other entities that control, are controlled by, or are under common
      control with that entity. For the purposes of this definition,
      "control" means (i) the power, direct or indirect, to cause the
      direction or management of such entity, whether by contract or
      otherwise, or (ii) ownership of fifty percent (50%) or more of the
      outstanding shares, or (iii) beneficial ownership of such entity.

      "You" (or "Your") shall mean an individual or Legal Entity
      exercising permissions granted by this License.

      "Source" form shall mean the preferred form for making modifications,
      including but not limited to software source code, documentation
      source, and configuration files.

      "Object" form shall mean any form resulting from mechanical
      transformation or translation of a Source form, including but
      not limited to compiled object code, generated documentation,
      and conversions to other media types.

      "Work" shall mean the work of authorship, whether in Source or
      Object form, made available under the License, as indicated by a
      copyright notice that is included in or attached to the work
      (an example is provided in the Appendix below).

      "Derivative Works" shall mean any work, whether in Source or Object
      form, that is based on (or derived from) the Work and for which the
      editorial revisions, annotations, elaborations, or other modifications
      represent, as a whole, an original work of authorship. For the purposes
      of this License, Derivative Works shall not include works that remain
      separable from, or merely link (or bind by name) to the interfaces of,
      the Work and Derivative Works thereof.

      "Contribution" shall mean any work of authorship, including
      the original version of the Work and any modifications or additions
      to that Work or Derivative Works thereof, that is intentionally
      submitted to Licensor for inclusion in the Work by the copyright owner
      or by an individual or Legal Entity authorized to submit on behalf of
      the copyright owner. For the purposes of this definition, "submitted"
      means any form of electronic, verbal, or written communication sent
      to the Licensor or its representatives, including but not limited to
      communication on electronic mailing lists, source code control systems,
      and issue tracking systems that are managed by, or on behalf of, the
      Licensor for the purpose of discussing and improving the Work, but
      excluding communication that is conspicuously marked or otherwise
      designated in writing by the copyright owner as "Not a Contribution."

      "Contributor" shall mean Licensor and any individual or Legal Entity
      on behalf of whom a Contribution has been received by Licensor and
      subsequently incorporated within the Work.

   2. Grant of Copyright License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      copyright license to reproduce, prepare Derivative Works of,
      publicly display, publicly perform, sublicense, and distribute the
      Work and such Derivative Works in Source or Object form.

   3. Grant of Patent License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      (except as stated in this section) patent license to make, have made,
      use, offer to sell, sell, import, and otherwise transfer the Work,
      where such license applies only to those patent claims licensable
      by such Contributor that are necessarily infringed by their
      Contribution(s) alone or by combination of their Contribution(s)
      with the Work to which such Contribution(s) was submitted. If You
      institute patent litigation against any entity (including a
      cross-claim or counterclaim in a lawsuit) alleging that the Work
      or a Contribution incorporated within the Work constitutes direct
      or contributory patent infringement, then any patent licenses
      granted to You under this License for that Work shall terminate
      as of the date such litigation is filed.

   4. Redistribution. You may reproduce and distribute copies of the
      Work or Derivative Works thereof in any medium, with or without
      modifications, and in Source or Object form, provided that You
      meet the following conditions:

      (a) You must give any other recipients of the Work or
          Derivative Works a copy of this License; and

      (b) You must cause any modified files to carry prominent notices
          stating that You changed the files; and

      (c) You must retain, in the Source form of any Derivative Works
          that You distribute, all copyright, patent, trademark, and
          attribution notices from the Source form of the Work,
          excluding those notices that do not pertain to any part of
          the Derivative Works; and

      (d) If the Work includes a "NOTICE" text file as part of its
          distribution, then any Derivative Works that You distribute must
          include a readable copy of the attribution notices contained
          within such NOTICE file, excluding those notices that do not
          pertain to any part of the Derivative Works, in at least one
          of the following places: within a NOTICE text file distributed
          as part of the Derivative Works; within the Source form or
          documentation, if provided along with the Derivative Works; or,
          within a display generated by the Derivative Works, if and
          wherever such third-party notices normally appear. The contents
          of the NOTICE file are for informational purposes only and
          do not modify the License. You may add Your own attribution
          notices within Derivative Works that You distribute, alongside
          or as an addendum to the NOTICE text from the Work, provided
          that such additional attribution notices cannot be construed
          as modifying the License.

      You may add Your own copyright statement to Your modifications and
      may provide additional or different license terms and conditions
      for use, reproduction, or distribution of Your modifications, or
      for any such Derivative Works as a whole, provided Your use,
      reproduction, and distribution of the Work otherwise complies with
      the conditions stated in this License.

   5. Submission of Contributions. Unless You explicitly state otherwise,
      any Contribution intentionally submitted for inclusion in the Work
      by You to the Licensor shall be under the terms and conditions of
      this License, without any additional terms or conditions.
      Notwithstanding the above, nothing herein shall supersede or modify
      the terms of any separate license agreement you may have executed
      with Licensor regarding such Contributions.

   6. Trademarks. This License does not grant permission to use the trade
      names, trademarks, service marks, or product names of the Licensor,
      except as required for reasonable and customary use in describing the
      origin of the Work and reproducing the content of the NOTICE file.

   7. Disclaimer of Warranty. Unless required by applicable law or
      agreed to in writing, Licensor provides the Work (and each
      Contributor provides its Contributions) on an "AS IS" BASIS,
      WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
      implied, including, without limitation, any warranties or conditions
      of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
      PARTICULAR PURPOSE. You are solely responsible for determining the
      appropriateness of using or redistributing the Work and assume any
      risks associated with Your exercise of permissions under this License.

   8. Limitation of Liability. In no event and under no legal theory,
      whether in tort (including negligence), contract, or otherwise,
      unless required by applicable law (such as deliberate and grossly
      negligent acts) or agreed to in writing, shall any Contributor be
      liable to You for damages, including any direct, indirect, special,
      incidental, or consequential damages of any character arising as a
      result of this License or out of the use or inability to use the
      Work (including but not limited to damages for loss of goodwill,
      work stoppage, computer failure or malfunction, or any and all
      other commercial damages or losses), even if such Contributor
      has been advised of the possibility of such damages.

   9. Accepting Warranty or Additional Liability. While redistributing
      the Work or Derivative Works thereof, You may choose to offer,
      and charge a fee for, acceptance of support, warranty, indemnity,
      or other liability obligations and/or rights consistent with this
      License. However, in accepting such obligations, You may act only
      on Your own behalf and on Your sole responsibility, not on behalf
      of any other Contributor, and only if You agree to indemnify,
      defend, and hold each Contributor harmless for any liability
      incurred by, or claims asserted against, such Contributor by reason
      of your accepting any such warranty or additional liability.

   END OF TERMS AND CONDITIONS

   APPENDIX: How to apply the Apache License to your work.

      To apply the Apache License to your work, attach the following
      boilerplate notice, with the fields enclosed by brackets "[]"
      replaced with your own identifying information. (Don't include
      the brackets!)  The text should be enclosed in the appropriate
      comment syntax for the file format. We also recommend that a
      file or class name and description of purpose be included on the
      same "printed page" as the copyright notice for easier
      identification within third-party archives.

   Copyright [yyyy] [name of copyright owner]

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"strings"
)

type TagPair struct {
	FreelingTag string
	SpacyTag    string
}

var TagPairs = []TagPair{
	{"NCMS000", "NOUN__Gender=Masc|Number=Sing"},
	{"NCMS00A", "NOUN__Degree=Aug|Gender=Masc|Number=Sing"},
	{"NCMS00D", "NOUN__Degree=Dim|Gender=Masc|Number=Sing"},
	{"NCMP000", "NOUN__Gender=Masc|Number=Plur"},
	{"NCMP00A", "NOUN__Degree=Aug|Gender=Masc|Number=Plur"},
	{"NCMP00D", "NOUN__Degree=Dim|Gender=Masc|Number=Plur"},
	{"NCFS000", "NOUN__Gender=Fem|Number=Sing"},
	{"NCFS00A", "NOUN__Degree=Aug|Gender=Fem|Number=Sing"},
	{"NCFS00D", "NOUN__Degree=Dim|Gender=Fem|Number=Sing"},
	{"NCFP000", "NOUN__Gender=Fem|Number=Plur"},
	{"NCFP00A", "NOUN__Degree=Aug|Gender=Fem|Number=Plur"},
	{"NCFP00D", "NOUN__Degree=Dim|Gender=Fem|Number=Plur"},
	{"NCCS000", "NOUN__Number=Sing"},
	{"NCCP000", "NOUN__Number=Plur"},
	{"NCNS000", "NOUN__Gender=Neut|Number=Sing"},
	{"NPMSS00", "PROPN__Gender=Masc|NameType=Prs|Number=Sing"},
	{"NPMSS0A", "PROPN__Degree=Aug|Gender=Masc|NameType=Prs|Number=Sing"},
	{"NPMSS0D", "PROPN__Degree=Dim|Gender=Masc|NameType=Prs|Number=Sing"},
	{"NPMPS00", "PROPN__Gender=Masc|NameType=Prs|Number=Plur"},
	{"NPMPS0A", "PROPN__Degree=Aug|Gender=Masc|NameType=Prs|Number=Plur"},
	{"NPMPS0D", "PROPN__Degree=Dim|Gender=Masc|NameType=Prs|Number=Plur"},
	{"NPFSS00", "PROPN__Gender=Fem|NameType=Prs|Number=Sing"},
	{"NPFSS0A", "PROPN__Degree=Aug|Gender=Fem|NameType=Prs|Number=Sing"},
	{"NPFSS0D", "PROPN__Degree=Dim|Gender=Fem|NameType=Prs|Number=Sing"},
	{"NPFPS00", "PROPN__Gender=Fem|NameType=Prs|Number=Plur"},
	{"NPFPS0A", "PROPN__Degree=Aug|Gender=Fem|NameType=Prs|Number=Plur"},
	{"NPFPS0D", "PROPN__Degree=Dim|Gender=Fem|NameType=Prs|Number=Plur"},
	{"NPCSS00", "PROPN__NameType=Prs|Number=Sing"},
	{"NPCPS00", "PROPN__NameType=Prs|Number=Plur"},
	{"NPMSG00", "PROPN__Gender=Masc|NameType=Geo|Number=Sing"},
	{"NPMPG00", "PROPN__Gender=Masc|NameType=Geo|Number=Plur"},
	{"NPFSG00", "PROPN__Gender=Fem|NameType=Geo|Number=Sing"},
	{"NPFPG00", "PROPN__Gender=Fem|NameType=Geo|Number=Plur"},
	{"NPCSG00", "PROPN__NameType=Geo|Number=Sing"},
	{"NPCPG00", "PROPN__NameType=Geo|Number=Plur"},
	{"NPMSO00", "PROPN__Gender=Masc|NameType=Com|Number=Sing"},
	{"NPMPO00", "PROPN__Gender=Masc|NameType=Com|Number=Plur"},
	{"NPFSO00", "PROPN__Gender=Fem|NameType=Com|Number=Sing"},
	{"NPFPO00", "PROPN__Gender=Fem|NameType=Com|Number=Plur"},
	{"NPCSO00", "PROPN__NameType=Com|Number=Sing"},
	{"NPCPO00", "PROPN__NameType=Com|Number=Plur"},
	{"NPMSV00", "PROPN__Gender=Masc|NameType=Oth|Number=Sing"},
	{"NPMPV00", "PROPN__Gender=Masc|NameType=Oth|Number=Plur"},
	{"NPFSV00", "PROPN__Gender=Fem|NameType=Oth|Number=Sing"},
	{"NPFPV00", "PROPN__Gender=Fem|NameType=Oth|Number=Plur"},
	{"NPCSV00", "PROPN__NameType=Oth|Number=Sing"},
	{"NPCPV00", "PROPN__NameType=Oth|Number=Plur"},
	{"AQVMS00", "ADJ__Degree=Pos|Gender=Masc|Number=Sing"},
	{"AQVMP00", "ADJ__Degree=Pos|Gender=Masc|Number=Plur"},
	{"AQVFS00", "ADJ__Degree=Pos|Gender=Fem|Number=Sing"},
	{"AQVFP00", "ADJ__Degree=Pos|Gender=Fem|Number=Plur"},
	{"AQVCS00", "ADJ__Degree=Pos|Number=Sing"},
	{"AQVCP00", "ADJ__Degree=Pos|Number=Plur"},
	{"AQVCN00", "ADJ__Degree=Pos"},
	{"AQSMS00", "ADJ__Degree=Sup|Gender=Masc|Number=Sing"},
	{"AQSMP00", "ADJ__Degree=Sup|Gender=Masc|Number=Plur"},
	{"AQSFS00", "ADJ__Degree=Sup|Gender=Fem|Number=Sing"},
	{"AQSFP00", "ADJ__Degree=Sup|Gender=Fem|Number=Plur"},
	{"AQSCS00", "ADJ__Degree=Sup|Number=Sing"},
	{"AQSCP00", "ADJ__Degree=Sup|Number=Plur"},
	{"AQSCN00", "ADJ__Degree=Sup"},
	{"APVMS1S", "ADJ__Degree=Pos|Gender=Masc|Number=Sing|Number[psor]=Sing|Person=1|Poss=Yes"},
	{"APVMS2S", "ADJ__Degree=Pos|Gender=Masc|Number=Sing|Number[psor]=Sing|Person=2|Poss=Yes"},
	{"APVMS3S", "ADJ__Degree=Pos|Gender=Masc|Number=Sing|Number[psor]=Sing|Person=3|Poss=Yes"},
	{"APVMS1P", "ADJ__Degree=Pos|Gender=Masc|Number=Sing|Number[psor]=Plur|Person=1|Poss=Yes"},
	{"APVMS2P", "ADJ__Degree=Pos|Gender=Masc|Number=Sing|Number[psor]=Plur|Person=2|Poss=Yes"},
	{"APVMS3P", "ADJ__Degree=Pos|Gender=Masc|Number=Sing|Number[psor]=Plur|Person=3|Poss=Yes"},
	{"APVMP1S", "ADJ__Degree=Pos|Gender=Masc|Number=Plur|Number[psor]=Sing|Person=1|Poss=Yes"},
	{"APVMP2S", "ADJ__Degree=Pos|Gender=Masc|Number=Plur|Number[psor]=Sing|Person=2|Poss=Yes"},
	{"APVMP3S", "ADJ__Degree=Pos|Gender=Masc|Number=Plur|Number[psor]=Sing|Person=3|Poss=Yes"},
	{"APVMP1P", "ADJ__Degree=Pos|Gender=Masc|Number=Plur|Number[psor]=Plur|Person=1|Poss=Yes"},
	{"APVMP2P", "ADJ__Degree=Pos|Gender=Masc|Number=Plur|Number[psor]=Plur|Person=2|Poss=Yes"},
	{"APVMP3P", "ADJ__Degree=Pos|Gender=Masc|Number=Plur|Number[psor]=Plur|Person=3|Poss=Yes"},
	{"APVFS1S", "ADJ__Degree=Pos|Gender=Fem|Number=Sing|Number[psor]=Sing|Person=1|Poss=Yes"},
	{"APVFS2S", "ADJ__Degree=Pos|Gender=Fem|Number=Sing|Number[psor]=Sing|Person=2|Poss=Yes"},
	{"APVFS3S", "ADJ__Degree=Pos|Gender=Fem|Number=Sing|Number[psor]=Sing|Person=3|Poss=Yes"},
	{"APVFS1P", "ADJ__Degree=Pos|Gender=Fem|Number=Sing|Number[psor]=Plur|Person=1|Poss=Yes"},
	{"APVFS2P", "ADJ__Degree=Pos|Gender=Fem|Number=Sing|Number[psor]=Plur|Person=2|Poss=Yes"},
	{"APVFS3P", "ADJ__Degree=Pos|Gender=Fem|Number=Sing|Number[psor]=Plur|Person=3|Poss=Yes"},
	{"APVFP1S", "ADJ__Degree=Pos|Gender=Fem|Number=Plur|Number[psor]=Sing|Person=1|Poss=Yes"},
	{"APVFP2S", "ADJ__Degree=Pos|Gender=Fem|Number=Plur|Number[psor]=Sing|Person=2|Poss=Yes"},
	{"APVFP3S", "ADJ__Degree=Pos|Gender=Fem|Number=Plur|Number[psor]=Sing|Person=3|Poss=Yes"},
	{"APVFP1P", "ADJ__Degree=Pos|Gender=Fem|Number=Plur|Number[psor]=Plur|Person=1|Poss=Yes"},
	{"APVFP2P", "ADJ__Degree=Pos|Gender=Fem|Number=Plur|Number[psor]=Plur|Person=2|Poss=Yes"},
	{"APVFP3P", "ADJ__Degree=Pos|Gender=Fem|Number=Plur|Number[psor]=Plur|Person=3|Poss=Yes"},
	{"APVCS1S", "ADJ__Degree=Pos|Number=Sing|Number[psor]=Sing|Person=1|Poss=Yes"},
	{"APVCS2S", "ADJ__Degree=Pos|Number=Sing|Number[psor]=Sing|Person=2|Poss=Yes"},
	{"APVCS3S", "ADJ__Degree=Pos|Number=Sing|Number[psor]=Sing|Person=3|Poss=Yes"},
	{"APVCS1P", "ADJ__Degree=Pos|Number=Sing|Number[psor]=Plur|Person=1|Poss=Yes"},
	{"APVCS2P", "ADJ__Degree=Pos|Number=Sing|Number[psor]=Plur|Person=2|Poss=Yes"},
	{"APVCS3P", "ADJ__Degree=Pos|Number=Sing|Number[psor]=Plur|Person=3|Poss=Yes"},
	{"APVCP1S", "ADJ__Degree=Pos|Number=Plur|Number[psor]=Sing|Person=1|Poss=Yes"},
	{"APVCP2S", "ADJ__Degree=Pos|Number=Plur|Number[psor]=Sing|Person=2|Poss=Yes"},
	{"APVCP3S", "ADJ__Degree=Pos|Number=Plur|Number[psor]=Sing|Person=3|Poss=Yes"},
	{"APVCP1P", "ADJ__Degree=Pos|Number=Plur|Number[psor]=Plur|Person=1|Poss=Yes"},
	{"APVCP2P", "ADJ__Degree=Pos|Number=Plur|Number[psor]=Plur|Person=2|Poss=Yes"},
	{"APVCP3P", "ADJ__Degree=Pos|Number=Plur|Number[psor]=Plur|Person=3|Poss=Yes"},
	{"APVCN1S", "ADJ__Degree=Pos|Number[psor]=Sing|Person=1|Poss=Yes"},
	{"APVCN2S", "ADJ__Degree=Pos|Number[psor]=Sing|Person=2|Poss=Yes"},
	{"APVCN3S", "ADJ__Degree=Pos|Number[psor]=Sing|Person=3|Poss=Yes"},
	{"APVCN1P", "ADJ__Degree=Pos|Number[psor]=Plur|Person=1|Poss=Yes"},
	{"APVCN2P", "ADJ__Degree=Pos|Number[psor]=Plur|Person=2|Poss=Yes"},
	{"APVCN3P", "ADJ__Degree=Pos|Number[psor]=Plur|Person=3|Poss=Yes"},
	{"AO0MS00", "ADJ__Gender=Masc|Number=Sing|NumType=Ord"},
	{"AO0MP00", "ADJ__Gender=Masc|Number=Plur|NumType=Ord"},
	{"AO0FS00", "ADJ__Gender=Fem|Number=Sing|NumType=Ord"},
	{"AO0FP00", "ADJ__Gender=Fem|Number=Plur|NumType=Ord"},
	{"AO0CS00", "ADJ__Number=Sing|NumType=Ord"},
	{"AO0CP00", "ADJ__Number=Plur|NumType=Ord"},
	{"AO0CN00", "ADJ__NumType=Ord"},
	{"PP10SN0", "PRON__Case=Nom|Number=Sing|Person=1|PronType=Prs"},
	{"PP10SD0", "PRON__Case=Dat|Number=Sing|Person=1|PronType=Prs"},
	{"PP10SA0", "PRON__Case=Acc|Number=Sing|Person=1|PronType=Prs"},
	{"PP20SN0", "PRON__Case=Nom|Number=Sing|Person=2|PronType=Prs"},
	{"PP20SD0", "PRON__Case=Dat|Number=Sing|Person=2|PronType=Prs"},
	{"PP20SA0", "PRON__Case=Acc|Number=Sing|Person=2|PronType=Prs"},
	{"PP20SNP", "PRON__Case=Nom|Number=Sing|Person=2|Polite=Form|PronType=Prs"},
	{"PP20SDP", "PRON__Case=Dat|Number=Sing|Person=2|Polite=Form|PronType=Prs"},
	{"PP20SAP", "PRON__Case=Acc|Number=Sing|Person=2|Polite=Form|PronType=Prs"},
	{"PP3MSN0", "PRON__Case=Nom|Gender=Masc|Number=Sing|Person=3|PronType=Prs"},
	{"PP3MSD0", "PRON__Case=Dat|Gender=Masc|Number=Sing|Person=3|PronType=Prs"},
	{"PP3MSA0", "PRON__Case=Acc|Gender=Masc|Number=Sing|Person=3|PronType=Prs"},
	{"PP3FSN0", "PRON__Case=Nom|Gender=Fem|Number=Sing|Person=3|PronType=Prs"},
	{"PP3FSD0", "PRON__Case=Dat|Gender=Fem|Number=Sing|Person=3|PronType=Prs"},
	{"PP3FSA0", "PRON__Case=Acc|Gender=Fem|Number=Sing|Person=3|PronType=Prs"},
	{"PP10PN0", "PRON__Case=Nom|Number=Plur|Person=1|PronType=Prs"},
	{"PP10PD0", "PRON__Case=Dat|Number=Plur|Person=1|PronType=Prs"},
	{"PP10PA0", "PRON__Case=Acc|Number=Plur|Person=1|PronType=Prs"},
	{"PP20PN0", "PRON__Case=Nom|Number=Plur|Person=2|PronType=Prs"},
	{"PP20PD0", "PRON__Case=Dat|Number=Plur|Person=2|PronType=Prs"},
	{"PP20PA0", "PRON__Case=Acc|Number=Plur|Person=2|PronType=Prs"},
	{"PP20PNP", "PRON__Case=Nom|Number=Plur|Person=2|Polite=Form|PronType=Prs"},
	{"PP20PDP", "PRON__Case=Dat|Number=Plur|Person=2|Polite=Form|PronType=Prs"},
	{"PP20PAP", "PRON__Case=Acc|Number=Plur|Person=2|Polite=Form|PronType=Prs"},
	{"PP3MPN0", "PRON__Case=Nom|Gender=Masc|Number=Plur|Person=3|PronType=Prs"},
	{"PP3MPD0", "PRON__Case=Dat|Gender=Masc|Number=Plur|Person=3|PronType=Prs"},
	{"PP3MPA0", "PRON__Case=Acc|Gender=Masc|Number=Plur|Person=3|PronType=Prs"},
	{"PP3FPN0", "PRON__Case=Nom|Gender=Fem|Number=Plur|Person=3|PronType=Prs"},
	{"PP3FPD0", "PRON__Case=Dat|Gender=Fem|Number=Plur|Person=3|PronType=Prs"},
	{"PP3FPA0", "PRON__Case=Acc|Gender=Fem|Number=Plur|Person=3|PronType=Prs"},
	{"PD00000", "PRON__PronType=Dem"},
	{"PT00000", "PRON__PronType=Int"},
	{"PR00000", "PRON__PronType=Rel"},
	{"PE00000", "PRON__PronType=Exc"},
	{"DA0MS0", "DET__Gender=Masc|Number=Sing|PronType=Art"},
	{"DA0FS0", "DET__Gender=Fem|Number=Sing|PronType=Art"},
	{"DA0MP0", "DET__Gender=Masc|Number=Plur|PronType=Art"},
	{"DA0FP0", "DET__Gender=Fem|Number=Plur|PronType=Art"},
	{"DP1MSS", "DET__Gender=Masc|Number=Sing|Number[psor]=Sing|Person=1|PronType=Prs"},
	{"DP1MPS", "DET__Gender=Masc|Number=Plur|Number[psor]=Sing|Person=1|PronType=Prs"},
	{"DP1FSS", "DET__Gender=Fem|Number=Sing|Number[psor]=Sing|Person=1|PronType=Prs"},
	{"DP1FPS", "DET__Gender=Fem|Number=Plur|Number[psor]=Sing|Person=1|PronType=Prs"},
	{"DP2MSS", "DET__Gender=Masc|Number=Sing|Number[psor]=Sing|Person=2|PronType=Prs"},
	{"DP2MPS", "DET__Gender=Masc|Number=Plur|Number[psor]=Sing|Person=2|PronType=Prs"},
	{"DP2FSS", "DET__Gender=Fem|Number=Sing|Number[psor]=Sing|Person=2|PronType=Prs"},
	{"DP2FPS", "DET__Gender=Fem|Number=Plur|Number[psor]=Sing|Person=2|PronType=Prs"},
	{"DP3MSS", "DET__Gender=Masc|Number=Sing|Number[psor]=Sing|Person=3|PronType=Prs"},
	{"DP3MPS", "DET__Gender=Masc|Number=Plur|Number[psor]=Sing|Person=3|PronType=Prs"},
	{"DP3FSS", "DET__Gender=Fem|Number=Sing|Number[psor]=Sing|Person=3|PronType=Prs"},
	{"DP3FPS", "DET__Gender=Fem|Number=Plur|Number[psor]=Sing|Person=3|PronType=Prs"},
	{"DP1MSP", "DET__Gender=Masc|Number=Sing|Number[psor]=Plur|Person=1|PronType=Prs"},
	{"DP1MPP", "DET__Gender=Masc|Number=Plur|Number[psor]=Plur|Person=1|PronType=Prs"},
	{"DP1FSP", "DET__Gender=Fem|Number=Sing|Number[psor]=Plur|Person=1|PronType=Prs"},
	{"DP1FPP", "DET__Gender=Fem|Number=Plur|Number[psor]=Plur|Person=1|PronType=Prs"},
	{"DP2MSP", "DET__Gender=Masc|Number=Sing|Number[psor]=Plur|Person=2|PronType=Prs"},
	{"DP2MPP", "DET__Gender=Masc|Number=Plur|Number[psor]=Plur|Person=2|PronType=Prs"},
	{"DP2FSP", "DET__Gender=Fem|Number=Sing|Number[psor]=Plur|Person=2|PronType=Prs"},
	{"DP2FPP", "DET__Gender=Fem|Number=Plur|Number[psor]=Plur|Person=2|PronType=Prs"},
	{"DP3MSP", "DET__Gender=Masc|Number=Sing|Number[psor]=Plur|Person=3|PronType=Prs"},
	{"DP3MPP", "DET__Gender=Masc|Number=Plur|Number[psor]=Plur|Person=3|PronType=Prs"},
	{"DP3FSP", "DET__Gender=Fem|Number=Sing|Number[psor]=Plur|Person=3|PronType=Prs"},
	{"DP3FPP", "DET__Gender=Fem|Number=Plur|Number[psor]=Plur|Person=3|PronType=Prs"},
	{"DD0MS0", "DET__Gender=Masc|Number=Sing|PronType=Dem"},
	{"DD0FS0", "DET__Gender=Fem|Number=Sing|PronType=Dem"},
	{"DD0MP0", "DET__Gender=Masc|Number=Plur|PronType=Dem"},
	{"DD0FP0", "DET__Gender=Fem|Number=Plur|PronType=Dem"},
	{"DT0000", "DET__PronType=Int"},
	{"DR0000", "DET__PronType=Rel"},
	{"DE0000", "DET__PronType=Exc"},
	{"Z", "NUM___"},
	{"W", "NUM___"},
	{"VMN0000", "VERB__VerbForm=Inf"},
	{"VMIP1S0", "VERB__Mood=Ind|Number=Sing|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VMIP2S0", "VERB__Mood=Ind|Number=Sing|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VMIP3S0", "VERB__Mood=Ind|Number=Sing|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VMIP1P0", "VERB__Mood=Ind|Number=Plur|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VMIP2P0", "VERB__Mood=Ind|Number=Plur|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VMIP3P0", "VERB__Mood=Ind|Number=Plur|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VMIF1S0", "VERB__Mood=Ind|Number=Sing|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VMIF2S0", "VERB__Mood=Ind|Number=Sing|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VMIF3S0", "VERB__Mood=Ind|Number=Sing|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VMIF1P0", "VERB__Mood=Ind|Number=Plur|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VMIF2P0", "VERB__Mood=Ind|Number=Plur|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VMIF3P0", "VERB__Mood=Ind|Number=Plur|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VMIS1S0", "VERB__Mood=Ind|Number=Sing|Person=1|Tense=Past|VerbForm=Fin"},
	{"VMIS2S0", "VERB__Mood=Ind|Number=Sing|Person=2|Tense=Past|VerbForm=Fin"},
	{"VMIS3S0", "VERB__Mood=Ind|Number=Sing|Person=3|Tense=Past|VerbForm=Fin"},
	{"VMIS1P0", "VERB__Mood=Ind|Number=Plur|Person=1|Tense=Past|VerbForm=Fin"},
	{"VMIS2P0", "VERB__Mood=Ind|Number=Plur|Person=2|Tense=Past|VerbForm=Fin"},
	{"VMIS3P0", "VERB__Mood=Ind|Number=Plur|Person=3|Tense=Past|VerbForm=Fin"},
	{"VMII1S0", "VERB__Mood=Ind|Number=Sing|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VMII2S0", "VERB__Mood=Ind|Number=Sing|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VMII3S0", "VERB__Mood=Ind|Number=Sing|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VMII1P0", "VERB__Mood=Ind|Number=Plur|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VMII2P0", "VERB__Mood=Ind|Number=Plur|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VMII3P0", "VERB__Mood=Ind|Number=Plur|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VMIM1S0", "VERB__Mood=Ind|Number=Sing|Person=1|Tense=Pqp|VerbForm=Fin"},
	{"VMIM2S0", "VERB__Mood=Ind|Number=Sing|Person=2|Tense=Pqp|VerbForm=Fin"},
	{"VMIM3S0", "VERB__Mood=Ind|Number=Sing|Person=3|Tense=Pqp|VerbForm=Fin"},
	{"VMIM1P0", "VERB__Mood=Ind|Number=Plur|Person=1|Tense=Pqp|VerbForm=Fin"},
	{"VMIM2P0", "VERB__Mood=Ind|Number=Plur|Person=2|Tense=Pqp|VerbForm=Fin"},
	{"VMIM3P0", "VERB__Mood=Ind|Number=Plur|Person=3|Tense=Pqp|VerbForm=Fin"},
	{"VMSP1S0", "VERB__Mood=Sub|Number=Sing|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VMSP2S0", "VERB__Mood=Sub|Number=Sing|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VMSP3S0", "VERB__Mood=Sub|Number=Sing|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VMSP1P0", "VERB__Mood=Sub|Number=Plur|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VMSP2P0", "VERB__Mood=Sub|Number=Plur|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VMSP3P0", "VERB__Mood=Sub|Number=Plur|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VMSF1S0", "VERB__Mood=Sub|Number=Sing|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VMSF2S0", "VERB__Mood=Sub|Number=Sing|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VMSF3S0", "VERB__Mood=Sub|Number=Sing|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VMSF1P0", "VERB__Mood=Sub|Number=Plur|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VMSF2P0", "VERB__Mood=Sub|Number=Plur|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VMSF3P0", "VERB__Mood=Sub|Number=Plur|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VMSI1S0", "VERB__Mood=Sub|Number=Sing|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VMSI2S0", "VERB__Mood=Sub|Number=Sing|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VMSI3S0", "VERB__Mood=Sub|Number=Sing|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VMSI1P0", "VERB__Mood=Sub|Number=Plur|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VMSI2P0", "VERB__Mood=Sub|Number=Plur|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VMSI3P0", "VERB__Mood=Sub|Number=Plur|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VM0C1S0", "VERB__Mood=Cnd|Number=Sing|Person=1"},
	{"VM0C2S0", "VERB__Mood=Cnd|Number=Sing|Person=2"},
	{"VM0C3S0", "VERB__Mood=Cnd|Number=Sing|Person=3"},
	{"VM0C1P0", "VERB__Mood=Cnd|Number=Plur|Person=1"},
	{"VM0C2P0", "VERB__Mood=Cnd|Number=Plur|Person=2"},
	{"VM0C3P0", "VERB__Mood=Cnd|Number=Plur|Person=3"},
	{"VMM02S0", "VERB__Mood=Imp|Number=Sing|Person=2|VerbForm=Fin"},
	{"VMM03S0", "VERB__Mood=Imp|Number=Sing|Person=3|VerbForm=Fin"},
	{"VMM01P0", "VERB__Mood=Imp|Number=Plur|Person=1|VerbForm=Fin"},
	{"VMM02P0", "VERB__Mood=Imp|Number=Plur|Person=2|VerbForm=Fin"},
	{"VMM03P0", "VERB__Mood=Imp|Number=Plur|Person=3|VerbForm=Fin"},
	{"VMPS0SM", "VERB__Gender=Masc|Number=Sing|Tense=Past|VerbForm=Part"},
	{"VMPS0SF", "VERB__Gender=Fem|Number=Sing|Tense=Past|VerbForm=Part"},
	{"VMPS0PM", "VERB__Gender=Masc|Number=Plur|Tense=Past|VerbForm=Part"},
	{"VMPS0PF", "VERB__Gender=Fem|Number=Plur|Tense=Past|VerbForm=Part"},
	{"VMGP0SM", "VERB__Gender=Masc|Number=Sing|Tense=Pres|VerbForm=Ger"},
	{"VMGP0SF", "VERB__Gender=Fem|Number=Sing|Tense=Pres|VerbForm=Ger"},
	{"VMGP0PM", "VERB__Gender=Masc|Number=Plur|Tense=Pres|VerbForm=Ger"},
	{"VMGP0PF", "VERB__Gender=Fem|Number=Plur|Tense=Pres|VerbForm=Ger"},
	{"VAN0000", "AUX__VerbForm=Inf"},
	{"VAIP1S0", "AUX__Mood=Ind|Number=Sing|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VAIP2S0", "AUX__Mood=Ind|Number=Sing|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VAIP3S0", "AUX__Mood=Ind|Number=Sing|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VAIP1P0", "AUX__Mood=Ind|Number=Plur|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VAIP2P0", "AUX__Mood=Ind|Number=Plur|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VAIP3P0", "AUX__Mood=Ind|Number=Plur|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VAIF1S0", "AUX__Mood=Ind|Number=Sing|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VAIF2S0", "AUX__Mood=Ind|Number=Sing|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VAIF3S0", "AUX__Mood=Ind|Number=Sing|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VAIF1P0", "AUX__Mood=Ind|Number=Plur|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VAIF2P0", "AUX__Mood=Ind|Number=Plur|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VAIF3P0", "AUX__Mood=Ind|Number=Plur|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VAIS1S0", "AUX__Mood=Ind|Number=Sing|Person=1|Tense=Past|VerbForm=Fin"},
	{"VAIS2S0", "AUX__Mood=Ind|Number=Sing|Person=2|Tense=Past|VerbForm=Fin"},
	{"VAIS3S0", "AUX__Mood=Ind|Number=Sing|Person=3|Tense=Past|VerbForm=Fin"},
	{"VAIS1P0", "AUX__Mood=Ind|Number=Plur|Person=1|Tense=Past|VerbForm=Fin"},
	{"VAIS2P0", "AUX__Mood=Ind|Number=Plur|Person=2|Tense=Past|VerbForm=Fin"},
	{"VAIS3P0", "AUX__Mood=Ind|Number=Plur|Person=3|Tense=Past|VerbForm=Fin"},
	{"VAII1S0", "AUX__Mood=Ind|Number=Sing|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VAII2S0", "AUX__Mood=Ind|Number=Sing|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VAII3S0", "AUX__Mood=Ind|Number=Sing|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VAII1P0", "AUX__Mood=Ind|Number=Plur|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VAII2P0", "AUX__Mood=Ind|Number=Plur|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VAII3P0", "AUX__Mood=Ind|Number=Plur|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VAIM1S0", "AUX__Mood=Ind|Number=Sing|Person=1|Tense=Pqp|VerbForm=Fin"},
	{"VAIM2S0", "AUX__Mood=Ind|Number=Sing|Person=2|Tense=Pqp|VerbForm=Fin"},
	{"VAIM3S0", "AUX__Mood=Ind|Number=Sing|Person=3|Tense=Pqp|VerbForm=Fin"},
	{"VAIM1P0", "AUX__Mood=Ind|Number=Plur|Person=1|Tense=Pqp|VerbForm=Fin"},
	{"VAIM2P0", "AUX__Mood=Ind|Number=Plur|Person=2|Tense=Pqp|VerbForm=Fin"},
	{"VAIM3P0", "AUX__Mood=Ind|Number=Plur|Person=3|Tense=Pqp|VerbForm=Fin"},
	{"VASP1S0", "AUX__Mood=Sub|Number=Sing|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VASP2S0", "AUX__Mood=Sub|Number=Sing|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VASP3S0", "AUX__Mood=Sub|Number=Sing|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VASP1P0", "AUX__Mood=Sub|Number=Plur|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VASP2P0", "AUX__Mood=Sub|Number=Plur|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VASP3P0", "AUX__Mood=Sub|Number=Plur|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VASF1S0", "AUX__Mood=Sub|Number=Sing|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VASF2S0", "AUX__Mood=Sub|Number=Sing|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VASF3S0", "AUX__Mood=Sub|Number=Sing|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VASF1P0", "AUX__Mood=Sub|Number=Plur|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VASF2P0", "AUX__Mood=Sub|Number=Plur|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VASF3P0", "AUX__Mood=Sub|Number=Plur|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VASI1S0", "AUX__Mood=Sub|Number=Sing|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VASI2S0", "AUX__Mood=Sub|Number=Sing|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VASI3S0", "AUX__Mood=Sub|Number=Sing|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VASI1P0", "AUX__Mood=Sub|Number=Plur|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VASI2P0", "AUX__Mood=Sub|Number=Plur|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VASI3P0", "AUX__Mood=Sub|Number=Plur|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VA0C1S0", "AUX__Mood=Cnd|Number=Sing|Person=1"},
	{"VA0C2S0", "AUX__Mood=Cnd|Number=Sing|Person=2"},
	{"VA0C3S0", "AUX__Mood=Cnd|Number=Sing|Person=3"},
	{"VA0C1P0", "AUX__Mood=Cnd|Number=Plur|Person=1"},
	{"VA0C2P0", "AUX__Mood=Cnd|Number=Plur|Person=2"},
	{"VA0C3P0", "AUX__Mood=Cnd|Number=Plur|Person=3"},
	{"VAM02S0", "AUX__Mood=Imp|Number=Sing|Person=2|VerbForm=Fin"},
	{"VAM03S0", "AUX__Mood=Imp|Number=Sing|Person=3|VerbForm=Fin"},
	{"VAM01P0", "AUX__Mood=Imp|Number=Plur|Person=1|VerbForm=Fin"},
	{"VAM02P0", "AUX__Mood=Imp|Number=Plur|Person=2|VerbForm=Fin"},
	{"VAM03P0", "AUX__Mood=Imp|Number=Plur|Person=3|VerbForm=Fin"},
	{"VAPS0SM", "AUX__Gender=Masc|Number=Sing|Tense=Past|VerbForm=Part"},
	{"VAPS0SF", "AUX__Gender=Fem|Number=Sing|Tense=Past|VerbForm=Part"},
	{"VAPS0PM", "AUX__Gender=Masc|Number=Plur|Tense=Past|VerbForm=Part"},
	{"VAPS0PF", "AUX__Gender=Fem|Number=Plur|Tense=Past|VerbForm=Part"},
	{"VAGP0SM", "AUX__Gender=Masc|Number=Sing|Tense=Pres|VerbForm=Ger"},
	{"VAGP0SF", "AUX__Gender=Fem|Number=Sing|Tense=Pres|VerbForm=Ger"},
	{"VAGP0PM", "AUX__Gender=Masc|Number=Plur|Tense=Pres|VerbForm=Ger"},
	{"VAGP0PF", "AUX__Gender=Fem|Number=Plur|Tense=Pres|VerbForm=Ger"},
	{"VSN0000", "AUX__VerbForm=Inf"},
	{"VSIP1S0", "AUX__Mood=Ind|Number=Sing|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VSIP2S0", "AUX__Mood=Ind|Number=Sing|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VSIP3S0", "AUX__Mood=Ind|Number=Sing|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VSIP1P0", "AUX__Mood=Ind|Number=Plur|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VSIP2P0", "AUX__Mood=Ind|Number=Plur|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VSIP3P0", "AUX__Mood=Ind|Number=Plur|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VSIF1S0", "AUX__Mood=Ind|Number=Sing|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VSIF2S0", "AUX__Mood=Ind|Number=Sing|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VSIF3S0", "AUX__Mood=Ind|Number=Sing|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VSIF1P0", "AUX__Mood=Ind|Number=Plur|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VSIF2P0", "AUX__Mood=Ind|Number=Plur|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VSIF3P0", "AUX__Mood=Ind|Number=Plur|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VSIS1S0", "AUX__Mood=Ind|Number=Sing|Person=1|Tense=Past|VerbForm=Fin"},
	{"VSIS2S0", "AUX__Mood=Ind|Number=Sing|Person=2|Tense=Past|VerbForm=Fin"},
	{"VSIS3S0", "AUX__Mood=Ind|Number=Sing|Person=3|Tense=Past|VerbForm=Fin"},
	{"VSIS1P0", "AUX__Mood=Ind|Number=Plur|Person=1|Tense=Past|VerbForm=Fin"},
	{"VSIS2P0", "AUX__Mood=Ind|Number=Plur|Person=2|Tense=Past|VerbForm=Fin"},
	{"VSIS3P0", "AUX__Mood=Ind|Number=Plur|Person=3|Tense=Past|VerbForm=Fin"},
	{"VSII1S0", "AUX__Mood=Ind|Number=Sing|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VSII2S0", "AUX__Mood=Ind|Number=Sing|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VSII3S0", "AUX__Mood=Ind|Number=Sing|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VSII1P0", "AUX__Mood=Ind|Number=Plur|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VSII2P0", "AUX__Mood=Ind|Number=Plur|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VSII3P0", "AUX__Mood=Ind|Number=Plur|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VSIM1S0", "AUX__Mood=Ind|Number=Sing|Person=1|Tense=Pqp|VerbForm=Fin"},
	{"VSIM2S0", "AUX__Mood=Ind|Number=Sing|Person=2|Tense=Pqp|VerbForm=Fin"},
	{"VSIM3S0", "AUX__Mood=Ind|Number=Sing|Person=3|Tense=Pqp|VerbForm=Fin"},
	{"VSIM1P0", "AUX__Mood=Ind|Number=Plur|Person=1|Tense=Pqp|VerbForm=Fin"},
	{"VSIM2P0", "AUX__Mood=Ind|Number=Plur|Person=2|Tense=Pqp|VerbForm=Fin"},
	{"VSIM3P0", "AUX__Mood=Ind|Number=Plur|Person=3|Tense=Pqp|VerbForm=Fin"},
	{"VSSP1S0", "AUX__Mood=Sub|Number=Sing|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VSSP2S0", "AUX__Mood=Sub|Number=Sing|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VSSP3S0", "AUX__Mood=Sub|Number=Sing|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VSSP1P0", "AUX__Mood=Sub|Number=Plur|Person=1|Tense=Pres|VerbForm=Fin"},
	{"VSSP2P0", "AUX__Mood=Sub|Number=Plur|Person=2|Tense=Pres|VerbForm=Fin"},
	{"VSSP3P0", "AUX__Mood=Sub|Number=Plur|Person=3|Tense=Pres|VerbForm=Fin"},
	{"VSSF1S0", "AUX__Mood=Sub|Number=Sing|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VSSF2S0", "AUX__Mood=Sub|Number=Sing|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VSSF3S0", "AUX__Mood=Sub|Number=Sing|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VSSF1P0", "AUX__Mood=Sub|Number=Plur|Person=1|Tense=Fut|VerbForm=Fin"},
	{"VSSF2P0", "AUX__Mood=Sub|Number=Plur|Person=2|Tense=Fut|VerbForm=Fin"},
	{"VSSF3P0", "AUX__Mood=Sub|Number=Plur|Person=3|Tense=Fut|VerbForm=Fin"},
	{"VSSI1S0", "AUX__Mood=Sub|Number=Sing|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VSSI2S0", "AUX__Mood=Sub|Number=Sing|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VSSI3S0", "AUX__Mood=Sub|Number=Sing|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VSSI1P0", "AUX__Mood=Sub|Number=Plur|Person=1|Tense=Imp|VerbForm=Fin"},
	{"VSSI2P0", "AUX__Mood=Sub|Number=Plur|Person=2|Tense=Imp|VerbForm=Fin"},
	{"VSSI3P0", "AUX__Mood=Sub|Number=Plur|Person=3|Tense=Imp|VerbForm=Fin"},
	{"VS0C1S0", "AUX__Mood=Cnd|Number=Sing|Person=1"},
	{"VS0C2S0", "AUX__Mood=Cnd|Number=Sing|Person=2"},
	{"VS0C3S0", "AUX__Mood=Cnd|Number=Sing|Person=3"},
	{"VS0C1P0", "AUX__Mood=Cnd|Number=Plur|Person=1"},
	{"VS0C2P0", "AUX__Mood=Cnd|Number=Plur|Person=2"},
	{"VS0C3P0", "AUX__Mood=Cnd|Number=Plur|Person=3"},
	{"VSM02S0", "AUX__Mood=Imp|Number=Sing|Person=2|VerbForm=Fin"},
	{"VSM03S0", "AUX__Mood=Imp|Number=Sing|Person=3|VerbForm=Fin"},
	{"VSM01P0", "AUX__Mood=Imp|Number=Plur|Person=1|VerbForm=Fin"},
	{"VSM02P0", "AUX__Mood=Imp|Number=Plur|Person=2|VerbForm=Fin"},
	{"VSM03P0", "AUX__Mood=Imp|Number=Plur|Person=3|VerbForm=Fin"},
	{"VSPS0SM", "AUX__Gender=Masc|Number=Sing|Tense=Past|VerbForm=Part"},
	{"VSPS0SF", "AUX__Gender=Fem|Number=Sing|Tense=Past|VerbForm=Part"},
	{"VSPS0PM", "AUX__Gender=Masc|Number=Plur|Tense=Past|VerbForm=Part"},
	{"VSPS0PF", "AUX__Gender=Fem|Number=Plur|Tense=Past|VerbForm=Part"},
	{"VSGP0SM", "AUX__Gender=Masc|Number=Sing|Tense=Pres|VerbForm=Ger"},
	{"VSGP0SF", "AUX__Gender=Fem|Number=Sing|Tense=Pres|VerbForm=Ger"},
	{"VSGP0PM", "AUX__Gender=Masc|Number=Plur|Tense=Pres|VerbForm=Ger"},
	{"VSGP0PF", "AUX__Gender=Fem|Number=Plur|Tense=Pres|VerbForm=Ger"},
	{"RG", "ADV___"},
	{"RN", "ADV__PronType=Neg"},
	{"SP", "ADP__AdpType=Prep"},
	{"CC", "CCONJ___"},
	{"CS", "SCONJ___"},
	{"I", "INTJ___"},
	{"Fd", "PUNCT__PunctType=Colo"},
	{"Fc", "PUNCT__PunctType=Comm"},
	{"Fs", "PUNCT___"},
	{"Faa", "PUNCT__PunctSide=Ini|PunctType=Excl"},
	{"Fat", "PUNCT__PunctSide=Fin|PunctType=Excl"},
	{"Fg", "PUNCT__PunctType=Dash"},
	{"Fz", "PUNCT___"},
	{"Ft", "PUNCT___"},
	{"Fp", "PUNCT__PunctType=Peri"},
	{"Fia", "PUNCT__PunctSide=Ini|PunctType=Qest"},
	{"Fit", "PUNCT__PunctSide=Fin|PunctType=Qest"},
	{"Fe", "PUNCT__PunctType=Quot"},
	{"Fra", "PUNCT__PunctSide=Ini|PunctType=Quot"},
	{"Frc", "PUNCT__PunctSide=Fin|PunctType=Quot"},
	{"Fx", "PUNCT__PunctType=Semi"},
	{"Fh", "PUNCT___"},
	{"Fpa", "PUNCT__PunctSide=Ini|PunctType=Brck"},
	{"Fpt", "PUNCT__PunctSide=Fin|PunctType=Brck"},
	{"Fca", "PUNCT__PunctSide=Ini|PunctType=Brck"},
	{"Fct", "PUNCT__PunctSide=Fin|PunctType=Brck"},
	{"Fla", "PUNCT__PunctSide=Ini|PunctType=Brck"},
	{"Flt", "PUNCT__PunctSide=Fin|PunctType=Brck"},

	// Added
	{"VMG0000", "VERB__Tense=Pres|VerbForm=Ger"},
}

type VerbTag struct {
	TagPair
	Features map[string]string
}

var allVerbTags = buildAllVerbTags()

func buildAllVerbTags() []VerbTag {
	verbTags := []VerbTag{}
	for _, tagPair := range TagPairs {
		tag := tagPair.SpacyTag
		if strings.HasPrefix(tag, "VERB__") {
			features := map[string]string{}
			for _, featurePair := range strings.Split(tag[6:len(tag)], "|") {
				values := strings.Split(featurePair, "=")
				features[values[0]] = values[1]
			}
			verbTag := VerbTag{TagPair: tagPair, Features: features}
			verbTags = append(verbTags, verbTag)
		}
	}
	return verbTags
}
