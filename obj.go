package main

////type FreeGroup struct {
////Id string `json:"id"`
////Version string `json:"version"`
////Configs map[string]*GroupConfig `json:"configs"`
////}
//
//type FreeConfig struct {
//Id string `json:"id"`
//Version string `json:"version"`
//Entries map[string]*string `json:"entries"`
//}
//
//type GroupConfig struct {
//Id string `json:"id"`
//Labels map[string]*string `json:"labels`
//Entries map[string]*string `json:"entries"`
//}
//
//GET("config/config_UUID/config_version") x4
//
//LIST("config/config_UUID/") => 4
//
//LIST("group/") =>
//	g1 v1 = g1v1l --------g1 v1
//		v2 = g1v2l			 v2
//		v3 = g1v3l			 v3
//	g2 v    g2v
//	g3 v    g3v
//(g1,g2,g3)
//
//LIST("group/group_id") =>
//	v1 2l = v1 l1--------  v1 l1
//		v1 l2				l2
//	v2 l  = v2 l           v2 l
//	v3 l  = v3 l           v3 l
//(v1, v2, v3)
//
//"config/config_UUID/config_version"
//
//"/group/group_id/group_version/labels/config/config_id"
//
//
//db = {
//[key                                    :                              ("entry1":"vrednost";"entry2:vrednost2";......)],
//[key                                    :                              ("entry1":"vrednost";"entry2:vrednost2";......)],
//[key                                    :                              ("entry1":"vrednost";"entry2:vrednost2";......)],
//[key                                    :                              ("entry1":"vrednost";"entry2:vrednost2";......)],
//[key                                    :                              ("entry1":"vrednost";"entry2:vrednost2";......)],
//[key                                    :                              ("entry1":"vrednost";"entry2:vrednost2";......)],
//[key                                    :                              ("entry1":"vrednost";"entry2:vrednost2";......)],
//[key                                    :                              ("entry1":"vrednost";"entry2:vrednost2";......)]
//}
//
