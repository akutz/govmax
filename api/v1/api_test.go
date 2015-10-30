package apiv1

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	//"time"
)

var smis *SMIS


var testingSID string

func init() {
	host := os.Getenv("GOVMAX_SMISHOST")
	port := os.Getenv("GOVMAX_SMISPORT")
	insecure, _ := strconv.ParseBool(os.Getenv("GOVMAX_INSECURE"))
	username := os.Getenv("GOVMAX_USERNAME")
	password := os.Getenv("GOVMAX_PASSWORD")

	var err error
	smis, err = New(host, port, insecure, username, password)
	if err != nil {
		panic(err)
	}
}


func TestGetStorageArrays(*testing.T) {

	arrays, err := smis.GetStorageArrays()
	//Setup array name for rest of tests
	testingSID = arrays.Entries[0].Content.I_ElementName
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v",arrays))
}

/*
func TestGetStoragePools(*testing.T) {

	pools, err := smis.GetStoragePools(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", pools))
}

func TestGetMaskingViews(*testing.T) {
        maskingViews, err := smis.GetMaskingViews(testingSID)
        if err != nil {
                panic(err)
        }

        fmt.Println(fmt.Sprintf("%+v",maskingViews))
}


func TestGetStorageGroups(*testing.T) {

	groups, err := smis.GetStorageGroups(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", groups))
}

func TestGetVolumes(*testing.T) {

	vols, err := smis.GetStorageVolumes(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols))
}

func TestPortGroups(*testing.T) {

	vols, err := smis.GetPortGroups(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols))
}

func TestInitiatorGroups(*testing.T) {

	vols, err := smis.GetHostGroups(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols))
}


func TestPostVolumes(*testing.T) {

	PostVolRequest := &PostVolumesReq{
		PostVolumesRequestContent : PostVolumesReqContent{
			AtType : "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageConfigurationService",
			ElementName : "test_vol",
			ElementType : "2",
			EMCNumberOfDevices : "1",
			Size : "123",
		},
	}
	queuedJob, err := smis.PostVolumes(PostVolRequest,testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v",queuedJob))
}

func TestPostCreateGroup(*testing.T) {
	curTime := time.Now()
	PostGroupRequest := &PostGroupReq{
		PostGroupRequestContent : PostGroupReqContent{
			AtType : "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerConfigurationService",
			SG_Name : "TestingSG_" + curTime.Format("Jan-2-2006--15-04-05"),
			Type : "4",
		},
	}
	storageGroup, err := smis.PostCreateGroup(PostGroupRequest,testingSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", storageGroup))
}

func TestGetStoragePoolSettings(*testing.T) {
	storagePools, err := smis.GetStoragePools(testingSID)
	if err != nil {
		panic(err)
	}

	storagePoolSettings, err := smis.GetStoragePoolSettings(storagePools.Entries[0].Content.I_ElementName,testingSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", storagePoolSettings))
}

func TestGetSLOs(*testing.T) {

    SLOs, err := smis.GetSLOs(testingSID)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%+v", SLOs)
}

func TestPostVolumeToSG(*testing.T) {

	PostVol2SGRequest := &PostVolumesToSGReq{
		PostVolumesToSGRequestContent : PostVolumesToSGReqContent{
			AtType : "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService",
			PostVolumesToSGRequestContentMG : PostVolumesToSGReqContentMG{
				AtType : "http://schemas.emc.com/ecom/edaa/root/emc/SE_DeviceMaskingGroup",
				//Change SMI_sg2 to any existing Storage Group ID
   		 		InstanceID : "SYMMETRIX-+-" + testingSID + "-+-SMI_sg2",
			},
			PostVolumesToSGRequestContentMember : []PostVolumesToSGReqContentMember{
				PostVolumesToSGReqContentMember {
				AtType : "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageVolume",
    			CreationClassName : "Symm_StorageVolume",
    			//Change DeviceID to existing Volume ID
    			DeviceID  : "00051",
    			SystemCreationClassName : "Symm_StorageSystem",
    			SystemName : "SYMMETRIX-+-" + testingSID,
    			},
    		},
		},
	}

	vol2SG, err := smis.PostVolumesToSG(PostVol2SGRequest,testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vol2SG))
}

*/

func TestPostStorageHardwareID(*testing.T) {
	PostSHIDRequest := &PostStorageHardwareIDReq{
		PostStorageHardwareIDRequestContent : PostStorageHardwareIDReqContent{
			AtType : "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageHardwareIDManagementService",
			IDType : "2",
			StorageID : "10000000c94e5d22",
		},
	}
	storageGroup, err := smis.PostStorageHardwareID(PostSHIDRequest,testingSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", storageGroup))
}

func TestPostInitiatorsToHG(*testing.T) {

	PostInit2SGRequest := &PostInitiatorsToHGReq{
		PostInitiatorsToHGRequestContent : PostInitiatorsToHGReqContent{
			AtType : "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService",
			PostInitiatorsToHGRequestContentMG : PostInitiatorsToHGReqContentMG{
				AtType : "http://schemas.emc.com/ecom/edaa/root/emc/SE_InitiatorMaskingGroup",
				//Change test_hg to any existing Host Group ID
   		 		InstanceID : "SYMMETRIX-+-" + testingSID + "-+-test_hg",
			},
			PostInitiatorsToHGRequestContentMember : []PostInitiatorsToHGReqContentMember{
				PostInitiatorsToHGReqContentMember {
				AtType : "http://schemas.emc.com/ecom/edaa/root/emc/SE_StorageHardwareID",
    			//Change InstanceID to existing Initiator 
    			InstanceID : "10000000C94E5D22",
    			},
    		},
		},
	}

	vol2SG, err := smis.PostInitiatorsToHG(PostInit2SGRequest,testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vol2SG))
}
/*

func TestPostPortsToPG(*testing.T) {

	PostPort2PGRequest := &PostPortsToPGReq{
		PostPortsToPGRequestContent : PostPortsToPGReqContent{
			AtType : "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService",
			PostPortsToPGRequestContentMG : PostPortsToPGReqContentMG{
				AtType : "http://schemas.emc.com/ecom/edaa/root/emc/SE_TargetMaskingGroup",
				//Change test_pg to any existing Port Group ID
   		 		InstanceID : "SYMMETRIX-+-" + testingSID + "-+-test_pg",
			},
			PostPortsToPGRequestContentMember : []PostPortsToPGReqContentMember{
				PostPortsToPGReqContentMember {
				AtType : "http://schemas.emc.com/ecom/edaa/root/emc/Symm_FCSCSIProtocolEndpoint",
    			CreationClassName : "Symm_FCSCSIProtocolEndpoint",
    			//Change Name to existing FE port 
    			Name : "5000097350159009",
    			SystemCreationClassName : "Symm_StorageProcessorSystem",
    			SystemName : "SYMMETRIX-+-" + testingSID + "-+-FA-1D-+-9",
    			},
    		},
		},
	}

	port2PG, err := smis.PostPortsToPG(PostPort2PGRequest,testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", port2PG))
}
*/