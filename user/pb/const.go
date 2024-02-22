package userpb

const (
	CLIENT_PLATFORM_IOS     = int64(1) // 客户端类型：iOS
	CLIENT_PLATFORM_ANDROID = int64(2) // 客户端类型：Android
	CLIENT_PLATFORM_PC      = int64(3) // 客户端类型：PC
)

type PlatformName string

const (
	IOOIKAAK_NAME  PlatformName = "之粮"
	IOOIKAAK2_NAME PlatformName = "之之粮"
	IOOIKAAK3_NAME PlatformName = "之之之粮"
)

const (
	APP_ID_ANDROID_IOOIKAAK     = "1003" // Android
	APP_ID_ANDROID_IOOIKAAK2    = "3002" // Android
	APP_ID_ANDROID_IOOIKAAK3    = "3003" // Android
	APP_ID_IOS_IOOIKAAK3        = "1001" // iOS
	APP_ID_IOS_RABBIT_IOOIKAAK3 = "1006" // iOS
	APP_ID_IOS_TANSUO_IOOIKAAK3 = "1009" // iOS
	APP_ID_IOS_IOOIKAAK_TANSUO  = "1010" // iOS
	PC_ID_IOOIKAAK              = "1008" // PC
	PC_ID_IOOIKAAK2             = "4002" // PC
	PC_ID_IOOIKAAK3             = "4003" // PC
)

var APP_NAMES = map[string]string{
	APP_ID_ANDROID_IOOIKAAK:     "安卓iooikaak",
	APP_ID_ANDROID_IOOIKAAK2:    "安卓iooikaak2",
	APP_ID_ANDROID_IOOIKAAK3:    "安卓iooikaak3",
	APP_ID_IOS_IOOIKAAK3:        "苹果iooikaak3",
	APP_ID_IOS_RABBIT_IOOIKAAK3: "苹果RabbitIooikaak3",
	APP_ID_IOS_IOOIKAAK_TANSUO:  "苹果iooikaak探索",
	APP_ID_IOS_TANSUO_IOOIKAAK3: "苹果iooikaak3",
	PC_ID_IOOIKAAK:              "PCiooikaak",
	PC_ID_IOOIKAAK2:             "PCiooikaak2",
	PC_ID_IOOIKAAK3:             "PCiooikaak3",
}

var COMMON_APP_NAMES = map[string]string{
	APP_ID_ANDROID_IOOIKAAK:     "iooikaak",
	APP_ID_ANDROID_IOOIKAAK2:    "iooikaak2",
	APP_ID_ANDROID_IOOIKAAK3:    "iooikaak3",
	APP_ID_IOS_IOOIKAAK3:        "iooikaak3",
	APP_ID_IOS_RABBIT_IOOIKAAK3: "rabbitIooikaak3",
	APP_ID_IOS_IOOIKAAK_TANSUO:  "iooikaak",
	APP_ID_IOS_TANSUO_IOOIKAAK3: "iooikaak3",
	PC_ID_IOOIKAAK:              "iooikaak",
	PC_ID_IOOIKAAK2:             "iooikaak2",
	PC_ID_IOOIKAAK3:             "iooikaak3",
}

var PLATFORM_NAMES = map[int64]string{
	CLIENT_PLATFORM_IOS:     "iOSApp",
	CLIENT_PLATFORM_ANDROID: "AndroidApp",
	CLIENT_PLATFORM_PC:      "PC",
}

func AppLaoyuegou(appID string) bool {
	return appID == APP_ID_ANDROID_IOOIKAAK || appID == APP_ID_IOS_IOOIKAAK3 || appID == PC_ID_IOOIKAAK
}

func AppTouxingmao(appID string) bool {
	return appID == APP_ID_ANDROID_IOOIKAAK3 || appID == APP_ID_IOS_TANSUO_IOOIKAAK3 || appID == PC_ID_IOOIKAAK3
}

func AppDuoduotu(appID string) bool {
	return appID == APP_ID_ANDROID_IOOIKAAK2 || appID == APP_ID_IOS_RABBIT_IOOIKAAK3 || appID == PC_ID_IOOIKAAK2 || appID == APP_ID_IOS_IOOIKAAK_TANSUO
}

func AppAndroid(appID string) bool {
	return appID == APP_ID_ANDROID_IOOIKAAK || appID == APP_ID_ANDROID_IOOIKAAK2 || appID == APP_ID_ANDROID_IOOIKAAK3
}

func AppiOS(appID string) bool {
	return appID == APP_ID_IOS_IOOIKAAK3 || appID == APP_ID_IOS_RABBIT_IOOIKAAK3 || appID == APP_ID_IOS_TANSUO_IOOIKAAK3 || appID == APP_ID_IOS_IOOIKAAK_TANSUO
}

func AppPC(appID string) bool {
	return appID == PC_ID_IOOIKAAK || appID == PC_ID_IOOIKAAK2 || appID == PC_ID_IOOIKAAK3
}

func CurrencyUnitByAppID(appID string) string {
	if AppDuoduotu(appID) {
		return string(IOOIKAAK2_NAME)
	} else if AppLaoyuegou(appID) {
		return string(IOOIKAAK_NAME)
	} else if AppTouxingmao(appID) {
		return string(IOOIKAAK3_NAME)
	}
	return string(IOOIKAAK_NAME)
}

const (
	NobilityLv1 = int64(1)
	NobilityLv2 = int64(2)
	NobilityLv3 = int64(3)
	NobilityLv4 = int64(4)
	NobilityLv5 = int64(5)
	NobilityLv6 = int64(6)
)

var Speeds = map[int64]float64{
	NobilityLv3: 0.05,
	NobilityLv4: 0.07,
	NobilityLv5: 0.1,
	NobilityLv6: 0.15,
}
