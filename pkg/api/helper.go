package api

import (
	"reflect"

	"google.golang.org/protobuf/proto"
)

// Request and response types from/to esphome
const (
	UndefinedTypeID = iota
	HelloRequestTypeID
	HelloResponseTypeID
	ConnectRequestTypeID
	ConnectResponseTypeID
	DisconnectRequestTypeID
	DisconnectResponseTypeID
	PingRequestTypeID
	PingResponseTypeID
	DeviceInfoRequestTypeID
	DeviceInfoResponseTypeID
	ListEntitiesRequestTypeID
	ListEntitiesBinarySensorResponseTypeID
	ListEntitiesCoverResponseTypeID
	ListEntitiesFanResponseTypeID
	ListEntitiesLightResponseTypeID
	ListEntitiesSensorResponseTypeID
	ListEntitiesSwitchResponseTypeID
	ListEntitiesTextSensorResponseTypeID
	ListEntitiesDoneResponseTypeID
	SubscribeStatesRequestTypeID
	BinarySensorStateResponseTypeID
	CoverStateResponseTypeID
	FanStateResponseTypeID
	LightStateResponseTypeID
	SensorStateResponseTypeID
	SwitchStateResponseTypeID
	TextSensorStateResponseTypeID
	SubscribeLogsRequestTypeID
	SubscribeLogsResponseTypeID
	CoverCommandRequestTypeID
	FanCommandRequestTypeID
	LightCommandRequestTypeID
	SwitchCommandRequestTypeID
	SubscribeHomeAssistantServicesRequestTypeID
	HomeAssistantServiceResponseTypeID
	GetTimeRequestTypeID
	GetTimeResponseTypeID
	SubscribeHomeAssistantStatesRequestTypeID
	SubscribeHomeAssistantStateResponseTypeID
	HomeAssistantStateResponseTypeID
	ListEntitiesServicesResponseTypeID
	ExecuteServiceRequestTypeID
	ListEntitiesCameraResponseTypeID
	CameraImageResponseTypeID
	CameraImageRequestTypeID
	ListEntitiesClimateResponseTypeID
	ClimateStateResponseTypeID
	ClimateCommandRequestTypeID
	ListEntitiesNumberResponseTypeID
	NumberStateResponseTypeID
	NumberCommandRequestTypeID
	ListEntitiesSelectResponseTypeID
	SelectStateResponseTypeID
	SelectCommandRequestTypeID
	UnknownTypeID55
	UnknownTypeID56
	UnknownTypeID57
	ListEntitiesLockResponseTypeID
	LockStateResponseTypeID
	LockCommandRequestTypeID
	ListEntitiesButtonResponseTypeID
	ButtonCommandRequestTypeID
	ListEntitiesMediaPlayerResponseTypeID
	MediaPlayerStateResponseTypeID
	MediaPlayerCommandRequestTypeID
	SubscribeBluetoothLEAdvertisementsRequestID
	BluetoothLEAdvertisementResponseID
	BluetoothDeviceRequestID
	BluetoothDeviceConnectionResponseID
	BluetoothGATTGetServicesRequestID
	BluetoothGATTGetServicesResponseID
	BluetoothGATTGetServicesDoneResponseID
	BluetoothGATTReadRequestID
	BluetoothGATTReadResponseID
	BluetoothGATTWriteRequestID
	BluetoothGATTReadDescriptorRequestID
	BluetoothGATTWriteDescriptorRequestID
	BluetoothGATTNotifyRequestID
	BluetoothGATTNotifyDataResponseID
	SubscribeBluetoothConnectionsFreeRequestID
	BluetoothConnectionsFreeResponseID
	BluetoothGATTErrorResponseID
	BluetoothGATTWriteResponseID
	BluetoothGATTNotifyResponseID
	BluetoothDevicePairingResponseID
	BluetoothDeviceUnpairingResponseID
	UnsubscribeBluetoothLEAdvertisementsRequestID
	BluetoothDeviceClearCacheResponseID
	SubscribeVoiceAssistantRequestID
	VoiceAssistantRequestID
	VoiceAssistantResponseID
	VoiceAssistantEventResponseID
	BluetoothLERawAdvertisementsResponseID
	ListEntitiesAlarmControlPanelResponseID
	AlarmControlPanelStateResponseID
	AlarmControlPanelCommandRequestID
	ListEntitiesTextResponseID
	TextStateResponseID
	TextCommandRequestID
	ListEntitiesDateResponseID
	DateStateResponseID
	DateCommandRequestID
	ListEntitiesTimeResponseID
	TimeStateResponseID
	TimeCommandRequestID
	VoiceAssistantAudioID
	ListEntitiesEventResponseID
	EventResponseID
	ListEntitiesValveResponseID
	ValveStateResponseID
	ValveCommandRequestID
	ListEntitiesDateTimeResponseID
	DateTimeStateResponseID
	DateTimeCommandRequestID
)

func TypeID(message interface{}) uint64 {
	if message == nil {
		return UndefinedTypeID
	}

	// convert from pointer to normal type
	if reflect.ValueOf(message).Kind() == reflect.Ptr {
		message = reflect.ValueOf(message).Elem().Interface()
	}
	switch message.(type) {
	case HelloRequest:
		return HelloRequestTypeID
	case HelloResponse:
		return HelloResponseTypeID
	case ConnectRequest:
		return ConnectRequestTypeID
	case ConnectResponse:
		return ConnectResponseTypeID
	case DisconnectRequest:
		return DisconnectRequestTypeID
	case DisconnectResponse:
		return DisconnectResponseTypeID
	case PingRequest:
		return PingRequestTypeID
	case PingResponse:
		return PingResponseTypeID
	case DeviceInfoRequest:
		return DeviceInfoRequestTypeID
	case DeviceInfoResponse:
		return DeviceInfoResponseTypeID
	case ListEntitiesRequest:
		return ListEntitiesRequestTypeID
	case ListEntitiesBinarySensorResponse:
		return ListEntitiesBinarySensorResponseTypeID
	case ListEntitiesCoverResponse:
		return ListEntitiesCoverResponseTypeID
	case ListEntitiesFanResponse:
		return ListEntitiesFanResponseTypeID
	case ListEntitiesLightResponse:
		return ListEntitiesLightResponseTypeID
	case ListEntitiesSensorResponse:
		return ListEntitiesSensorResponseTypeID
	case ListEntitiesSwitchResponse:
		return ListEntitiesSwitchResponseTypeID
	case ListEntitiesTextSensorResponse:
		return ListEntitiesTextSensorResponseTypeID
	case ListEntitiesDoneResponse:
		return ListEntitiesDoneResponseTypeID
	case SubscribeStatesRequest:
		return SubscribeStatesRequestTypeID
	case BinarySensorStateResponse:
		return BinarySensorStateResponseTypeID
	case CoverStateResponse:
		return CoverStateResponseTypeID
	case FanStateResponse:
		return FanStateResponseTypeID
	case LightStateResponse:
		return LightStateResponseTypeID
	case SensorStateResponse:
		return SensorStateResponseTypeID
	case SwitchStateResponse:
		return SwitchStateResponseTypeID
	case TextSensorStateResponse:
		return TextSensorStateResponseTypeID
	case SubscribeLogsRequest:
		return SubscribeLogsRequestTypeID
	case SubscribeLogsResponse:
		return SubscribeLogsResponseTypeID
	case CoverCommandRequest:
		return CoverCommandRequestTypeID
	case FanCommandRequest:
		return FanCommandRequestTypeID
	case LightCommandRequest:
		return LightCommandRequestTypeID
	case SwitchCommandRequest:
		return SwitchCommandRequestTypeID
	case SubscribeHomeassistantServicesRequest:
		return SubscribeHomeAssistantServicesRequestTypeID
	case HomeassistantServiceResponse:
		return HomeAssistantServiceResponseTypeID
	case GetTimeRequest:
		return GetTimeRequestTypeID
	case GetTimeResponse:
		return GetTimeResponseTypeID
	case SubscribeHomeAssistantStatesRequest:
		return SubscribeHomeAssistantStatesRequestTypeID
	case SubscribeHomeAssistantStateResponse:
		return SubscribeHomeAssistantStateResponseTypeID
	case HomeAssistantStateResponse:
		return HomeAssistantStateResponseTypeID
	case ListEntitiesServicesResponse:
		return ListEntitiesServicesResponseTypeID
	case ExecuteServiceRequest:
		return ExecuteServiceRequestTypeID
	case ListEntitiesCameraResponse:
		return ListEntitiesCameraResponseTypeID
	case CameraImageResponse:
		return CameraImageResponseTypeID
	case CameraImageRequest:
		return CameraImageRequestTypeID
	case ListEntitiesClimateResponse:
		return ListEntitiesClimateResponseTypeID
	case ClimateStateResponse:
		return ClimateStateResponseTypeID
	case ClimateCommandRequest:
		return ClimateCommandRequestTypeID
	case ListEntitiesNumberResponse:
		return ListEntitiesNumberResponseTypeID
	case NumberStateResponse:
		return NumberStateResponseTypeID
	case NumberCommandRequest:
		return NumberCommandRequestTypeID
	case ListEntitiesSelectResponse:
		return ListEntitiesSelectResponseTypeID
	case SelectStateResponse:
		return SelectStateResponseTypeID
	case SelectCommandRequest:
		return SelectCommandRequestTypeID
	case ListEntitiesLockResponse:
		return ListEntitiesLockResponseTypeID
	case LockStateResponse:
		return LockStateResponseTypeID
	case LockCommandRequest:
		return LockCommandRequestTypeID
	case ListEntitiesButtonResponse:
		return ListEntitiesButtonResponseTypeID
	case ButtonCommandRequest:
		return ButtonCommandRequestTypeID
	case ListEntitiesMediaPlayerResponse:
		return ListEntitiesMediaPlayerResponseTypeID
	case MediaPlayerStateResponse:
		return MediaPlayerStateResponseTypeID
	case MediaPlayerCommandRequest:
		return MediaPlayerCommandRequestTypeID
	case SubscribeBluetoothLEAdvertisementsRequest:
		return SubscribeBluetoothLEAdvertisementsRequestID
	case BluetoothLEAdvertisementResponse:
		return BluetoothLEAdvertisementResponseID
	case BluetoothDeviceRequest:
		return BluetoothDeviceRequestID
	case BluetoothDeviceConnectionResponse:
		return BluetoothDeviceConnectionResponseID
	case BluetoothGATTGetServicesRequest:
		return BluetoothGATTGetServicesRequestID
	case BluetoothGATTGetServicesResponse:
		return BluetoothGATTGetServicesResponseID
	case BluetoothGATTGetServicesDoneResponse:
		return BluetoothGATTGetServicesDoneResponseID
	case BluetoothGATTReadRequest:
		return BluetoothGATTReadRequestID
	case BluetoothGATTReadResponse:
		return BluetoothGATTReadResponseID
	case BluetoothGATTWriteRequest:
		return BluetoothGATTWriteRequestID
	case BluetoothGATTReadDescriptorRequest:
		return BluetoothGATTReadDescriptorRequestID
	case BluetoothGATTWriteDescriptorRequest:
		return BluetoothGATTWriteDescriptorRequestID
	case BluetoothGATTNotifyRequest:
		return BluetoothGATTNotifyRequestID
	case BluetoothGATTNotifyDataResponse:
		return BluetoothGATTNotifyDataResponseID
	case SubscribeBluetoothConnectionsFreeRequest:
		return SubscribeBluetoothConnectionsFreeRequestID
	case BluetoothConnectionsFreeResponse:
		return BluetoothConnectionsFreeResponseID
	case BluetoothGATTErrorResponse:
		return BluetoothGATTErrorResponseID
	case BluetoothGATTWriteResponse:
		return BluetoothGATTWriteResponseID
	case BluetoothGATTNotifyResponse:
		return BluetoothGATTNotifyResponseID
	case BluetoothDevicePairingResponse:
		return BluetoothDevicePairingResponseID
	case BluetoothDeviceUnpairingResponse:
		return BluetoothDeviceUnpairingResponseID
	case UnsubscribeBluetoothLEAdvertisementsRequest:
		return UnsubscribeBluetoothLEAdvertisementsRequestID
	case BluetoothDeviceClearCacheResponse:
		return BluetoothDeviceClearCacheResponseID
	case SubscribeVoiceAssistantRequest:
		return SubscribeVoiceAssistantRequestID
	case VoiceAssistantRequest:
		return VoiceAssistantRequestID
	case VoiceAssistantResponse:
		return VoiceAssistantResponseID
	case VoiceAssistantEventResponse:
		return VoiceAssistantEventResponseID
	case BluetoothLERawAdvertisementsResponse:
		return BluetoothLERawAdvertisementsResponseID
	case ListEntitiesAlarmControlPanelResponse:
		return ListEntitiesAlarmControlPanelResponseID
	case AlarmControlPanelStateResponse:
		return AlarmControlPanelStateResponseID
	case AlarmControlPanelCommandRequest:
		return AlarmControlPanelCommandRequestID
	case ListEntitiesTextResponse:
		return ListEntitiesTextResponseID
	case TextStateResponse:
		return TextStateResponseID
	case TextCommandRequest:
		return TextCommandRequestID
	case ListEntitiesDateResponse:
		return ListEntitiesDateResponseID
	case DateStateResponse:
		return DateStateResponseID
	case DateCommandRequest:
		return DateCommandRequestID
	case ListEntitiesTimeResponse:
		return ListEntitiesTimeResponseID
	case TimeStateResponse:
		return TimeStateResponseID
	case TimeCommandRequest:
		return TimeCommandRequestID
	case VoiceAssistantAudio:
		return VoiceAssistantAudioID
	case ListEntitiesEventResponse:
		return ListEntitiesEventResponseID
	case EventResponse:
		return EventResponseID
	case ListEntitiesValveResponse:
		return ListEntitiesValveResponseID
	case ValveStateResponse:
		return ValveStateResponseID
	case ValveCommandRequest:
		return ValveCommandRequestID
	case ListEntitiesDateTimeResponse:
		return ListEntitiesDateTimeResponseID
	case DateTimeStateResponse:
		return DateTimeStateResponseID
	default:
		return UndefinedTypeID
	}
}

func NewMessageByTypeID(typeID uint64) proto.Message {
	switch typeID {
	case HelloRequestTypeID:
		return new(HelloRequest)
	case HelloResponseTypeID:
		return new(HelloResponse)
	case ConnectRequestTypeID:
		return new(ConnectRequest)
	case ConnectResponseTypeID:
		return new(ConnectResponse)
	case DisconnectRequestTypeID:
		return new(DisconnectRequest)
	case DisconnectResponseTypeID:
		return new(DisconnectResponse)
	case PingRequestTypeID:
		return new(PingRequest)
	case PingResponseTypeID:
		return new(PingResponse)
	case DeviceInfoRequestTypeID:
		return new(DeviceInfoRequest)
	case DeviceInfoResponseTypeID:
		return new(DeviceInfoResponse)
	case ListEntitiesRequestTypeID:
		return new(ListEntitiesRequest)
	case ListEntitiesBinarySensorResponseTypeID:
		return new(ListEntitiesBinarySensorResponse)
	case ListEntitiesCoverResponseTypeID:
		return new(ListEntitiesCoverResponse)
	case ListEntitiesFanResponseTypeID:
		return new(ListEntitiesFanResponse)
	case ListEntitiesLightResponseTypeID:
		return new(ListEntitiesLightResponse)
	case ListEntitiesSensorResponseTypeID:
		return new(ListEntitiesSensorResponse)
	case ListEntitiesSwitchResponseTypeID:
		return new(ListEntitiesSwitchResponse)
	case ListEntitiesTextSensorResponseTypeID:
		return new(ListEntitiesTextSensorResponse)
	case ListEntitiesDoneResponseTypeID:
		return new(ListEntitiesDoneResponse)
	case SubscribeStatesRequestTypeID:
		return new(SubscribeStatesRequest)
	case BinarySensorStateResponseTypeID:
		return new(BinarySensorStateResponse)
	case CoverStateResponseTypeID:
		return new(CoverStateResponse)
	case FanStateResponseTypeID:
		return new(FanStateResponse)
	case LightStateResponseTypeID:
		return new(LightStateResponse)
	case SensorStateResponseTypeID:
		return new(SensorStateResponse)
	case SwitchStateResponseTypeID:
		return new(SwitchStateResponse)
	case TextSensorStateResponseTypeID:
		return new(TextSensorStateResponse)
	case SubscribeLogsRequestTypeID:
		return new(SubscribeLogsRequest)
	case SubscribeLogsResponseTypeID:
		return new(SubscribeLogsResponse)
	case CoverCommandRequestTypeID:
		return new(CoverCommandRequest)
	case FanCommandRequestTypeID:
		return new(FanCommandRequest)
	case LightCommandRequestTypeID:
		return new(LightCommandRequest)
	case SwitchCommandRequestTypeID:
		return new(SwitchCommandRequest)
	case SubscribeHomeAssistantServicesRequestTypeID:
		return new(SubscribeHomeassistantServicesRequest)
	case HomeAssistantServiceResponseTypeID:
		return new(HomeassistantServiceResponse)
	case GetTimeRequestTypeID:
		return new(GetTimeRequest)
	case GetTimeResponseTypeID:
		return new(GetTimeResponse)
	case SubscribeHomeAssistantStatesRequestTypeID:
		return new(SubscribeHomeAssistantStatesRequest)
	case SubscribeHomeAssistantStateResponseTypeID:
		return new(SubscribeHomeAssistantStateResponse)
	case HomeAssistantStateResponseTypeID:
		return new(HomeAssistantStateResponse)
	case ListEntitiesServicesResponseTypeID:
		return new(ListEntitiesServicesResponse)
	case ExecuteServiceRequestTypeID:
		return new(ExecuteServiceRequest)
	case ListEntitiesCameraResponseTypeID:
		return new(ListEntitiesCameraResponse)
	case CameraImageResponseTypeID:
		return new(CameraImageResponse)
	case CameraImageRequestTypeID:
		return new(CameraImageRequest)
	case ListEntitiesClimateResponseTypeID:
		return new(ListEntitiesClimateResponse)
	case ClimateStateResponseTypeID:
		return new(ClimateStateResponse)
	case ClimateCommandRequestTypeID:
		return new(ClimateCommandRequest)
	case ListEntitiesNumberResponseTypeID:
		return new(ListEntitiesNumberResponse)
	case NumberStateResponseTypeID:
		return new(NumberStateResponse)
	case NumberCommandRequestTypeID:
		return new(NumberCommandRequest)
	case ListEntitiesSelectResponseTypeID:
		return new(ListEntitiesSelectResponse)
	case SelectStateResponseTypeID:
		return new(SelectStateResponse)
	case SelectCommandRequestTypeID:
		return new(SelectCommandRequest)
	case ListEntitiesLockResponseTypeID:
		return new(ListEntitiesLockResponse)
	case LockStateResponseTypeID:
		return new(LockStateResponse)
	case LockCommandRequestTypeID:
		return new(LockCommandRequest)
	case ListEntitiesButtonResponseTypeID:
		return new(ListEntitiesButtonResponse)
	case ButtonCommandRequestTypeID:
		return new(ButtonCommandRequest)
	case ListEntitiesMediaPlayerResponseTypeID:
		return new(ListEntitiesMediaPlayerResponse)
	case MediaPlayerStateResponseTypeID:
		return new(MediaPlayerStateResponse)
	case MediaPlayerCommandRequestTypeID:
		return new(MediaPlayerCommandRequest)
	case SubscribeBluetoothLEAdvertisementsRequestID:
		return new(SubscribeBluetoothLEAdvertisementsRequest)
	case BluetoothLEAdvertisementResponseID:
		return new(BluetoothLEAdvertisementResponse)
	case BluetoothDeviceRequestID:
		return new(BluetoothDeviceRequest)
	case BluetoothDeviceConnectionResponseID:
		return new(BluetoothDeviceConnectionResponse)
	case BluetoothGATTGetServicesRequestID:
		return new(BluetoothGATTGetServicesRequest)
	case BluetoothGATTGetServicesResponseID:
		return new(BluetoothGATTGetServicesResponse)
	case BluetoothGATTGetServicesDoneResponseID:
		return new(BluetoothGATTGetServicesDoneResponse)
	case BluetoothGATTReadRequestID:
		return new(BluetoothGATTReadRequest)
	case BluetoothGATTReadResponseID:
		return new(BluetoothGATTReadResponse)
	case BluetoothGATTWriteRequestID:
		return new(BluetoothGATTWriteRequest)
	case BluetoothGATTReadDescriptorRequestID:
		return new(BluetoothGATTReadDescriptorRequest)
	case BluetoothGATTWriteDescriptorRequestID:
		return new(BluetoothGATTWriteDescriptorRequest)
	case BluetoothGATTNotifyRequestID:
		return new(BluetoothGATTNotifyRequest)
	case BluetoothGATTNotifyDataResponseID:
		return new(BluetoothGATTNotifyDataResponse)
	case SubscribeBluetoothConnectionsFreeRequestID:
		return new(SubscribeBluetoothConnectionsFreeRequest)
	case BluetoothConnectionsFreeResponseID:
		return new(BluetoothConnectionsFreeResponse)
	case BluetoothGATTErrorResponseID:
		return new(BluetoothGATTErrorResponse)
	case BluetoothGATTWriteResponseID:
		return new(BluetoothGATTWriteResponse)
	case BluetoothGATTNotifyResponseID:
		return new(BluetoothGATTNotifyResponse)
	case BluetoothDevicePairingResponseID:
		return new(BluetoothDevicePairingResponse)
	case BluetoothDeviceUnpairingResponseID:
		return new(BluetoothDeviceUnpairingResponse)
	case UnsubscribeBluetoothLEAdvertisementsRequestID:
		return new(UnsubscribeBluetoothLEAdvertisementsRequest)
	case BluetoothDeviceClearCacheResponseID:
		return new(BluetoothDeviceClearCacheResponse)
	case SubscribeVoiceAssistantRequestID:
		return new(SubscribeVoiceAssistantRequest)
	case VoiceAssistantRequestID:
		return new(VoiceAssistantRequest)
	case VoiceAssistantResponseID:
		return new(VoiceAssistantResponse)
	case VoiceAssistantEventResponseID:
		return new(VoiceAssistantEventResponse)
	case BluetoothLERawAdvertisementsResponseID:
		return new(BluetoothLERawAdvertisementsResponse)
	case ListEntitiesAlarmControlPanelResponseID:
		return new(ListEntitiesAlarmControlPanelResponse)
	case AlarmControlPanelStateResponseID:
		return new(AlarmControlPanelStateResponse)
	case AlarmControlPanelCommandRequestID:
		return new(AlarmControlPanelCommandRequest)
	case ListEntitiesTextResponseID:
		return new(ListEntitiesTextResponse)
	case TextStateResponseID:
		return new(TextStateResponse)
	case TextCommandRequestID:
		return new(TextCommandRequest)
	case ListEntitiesDateResponseID:
		return new(ListEntitiesDateResponse)
	case DateStateResponseID:
		return new(DateStateResponse)
	case DateCommandRequestID:
		return new(DateCommandRequest)
	case ListEntitiesTimeResponseID:
		return new(ListEntitiesTimeResponse)
	case TimeStateResponseID:
		return new(TimeStateResponse)
	case TimeCommandRequestID:
		return new(TimeCommandRequest)
	case VoiceAssistantAudioID:
		return new(VoiceAssistantAudio)
	case ListEntitiesEventResponseID:
		return new(ListEntitiesEventResponse)
	case EventResponseID:
		return new(EventResponse)
	case ListEntitiesValveResponseID:
		return new(ListEntitiesValveResponse)
	case ValveStateResponseID:
		return new(ValveStateResponse)
	case ValveCommandRequestID:
		return new(ValveCommandRequest)
	case ListEntitiesDateTimeResponseID:
		return new(ListEntitiesDateTimeResponse)
	case DateTimeStateResponseID:
		return new(DateTimeStateResponse)
	case DateTimeCommandRequestID:
		return new(DateTimeCommandRequest)
	default:
		return nil
	}
}
