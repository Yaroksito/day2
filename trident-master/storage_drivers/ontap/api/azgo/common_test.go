package azgo

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateZAPIResponse(t *testing.T) {
	// Sample of exploit-capable xml
	// from https://github.com/mattermost/xml-roundtrip-validator/blob/master/validator_test.go used as a sanity test
	simpleInvalidResponseBody := `<Root ::attr="x">]]><x::Element/></Root>`

	sampleResponse := &http.Response{
		Body:          io.NopCloser(bytes.NewBufferString(simpleInvalidResponseBody)),
		ContentLength: int64(len(simpleInvalidResponseBody)),
	}
	_, err := ValidateZAPIResponse(sampleResponse)
	assert.Error(t, err, "should error on input with colons in an attribute's name")

	// <! <<!-- -->!-->"--> " ><! ">" <X/>>
	// Is an invalid xml directive
	// Make sure that the validator combined with Go 1.17+ passes the invalid directive in a ZAPI XML response
	// that would be valid in versions of Trident without the validation or would fail validation with older Go versions.
	validToUnpatchedTridentResponseBody := `<?xml version='1.0' encoding='UTF-8' ?>
<!DOCTYPE netapp SYSTEM 'file:/etc/netapp_gx.dtd'><netapp version='1.180' xmlns='http://www.netapp.com/filer/admin'>
<! <<!-- -->!-->"--> " ><! ">" <X/>><results status="passed"><attributes-list><initiator-group-info>
<initiator-group-alua-enabled>true</initiator-group-alua-enabled><initiator-group-delete-on-unmap>false
</initiator-group-delete-on-unmap><initiator-group-name>presnellbr</initiator-group-name>
<initiator-group-os-type>linux</initiator-group-os-type><initiator-group-throttle-borrow>false
</initiator-group-throttle-borrow><initiator-group-throttle-reserve>0</initiator-group-throttle-reserve>
<initiator-group-type>iscsi</initiator-group-type><initiator-group-use-partner>true</initiator-group-use-partner>
<initiator-group-uuid>1c7b3edc-aead-11eb-b8e0-00a0986e75a0</initiator-group-uuid>
<initiator-group-vsa-enabled>false</initiator-group-vsa-enabled><initiators><initiator-info>
<initiator-name>iqn.2005-03.org.open-iscsi:cb50e2753efd</initiator-name>
</initiator-info></initiators><vserver>CXE</vserver></initiator-group-info></attributes-list><num-records>1
</num-records></results></netapp>"`

	sampleResponse = &http.Response{
		Body:          io.NopCloser(bytes.NewBufferString(validToUnpatchedTridentResponseBody)),
		ContentLength: int64(len(validToUnpatchedTridentResponseBody)),
	}

	_, err = ValidateZAPIResponse(sampleResponse)
	assert.NoError(t, err, "an invalid XML directive in an otherwise valid ZAPI XML response caused"+
		" the validation to return an error")

	var zapiResp IgroupGetIterResponse
	unmarshalErr := xml.Unmarshal([]byte(validToUnpatchedTridentResponseBody), &zapiResp)
	errMsg := "the xml document should be unmarshalled without issue when not validated properly"
	assert.Nil(t, unmarshalErr, errMsg)
}
