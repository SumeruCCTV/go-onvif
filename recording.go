package onvif

func (device Device) GetRecordingConfiguration(recordingToken string) (interface{}, error) {
	// create soap
	soap := SOAP{
		User:     device.User,
		Password: device.Password,
		Body: `<GetRecordingConfiguration xmlns="http://www.onvif.org/ver10/recording/wsdl">
						<RecordingToken>` + recordingToken + `</RecordingToken>
					</GetRecordingConfiguration>`,
	}

	// send request
	response, err := soap.SendRequest(device.XAddr)
	if err != nil {
		return nil, err
	}

	// parse response
	result, err := response.ValueForPath("Envelope.Body.GetRecordingConfigurationResponse")
	if err != nil {
		return result, err
	}
	return result, nil
}
