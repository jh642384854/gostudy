package model

type User struct {
	Age                      int      `json:"age"`
	AvatarPhotoID            int64    `json:"avatarphotoid"`
	AvatarPraiseCount        int      `json:"avatarpraisecount"`
	AvatarPraised            bool     `json:"avatarpraised"`
	AvatarURL                string   `json:"avatarurl"`
	BasicInfo                []string `json:"basicinfo"`
	DetailInfo               []string `json:"detailinfo"`
	EducationString          string   `json:"educationstring"`
	EmotionStatus            int      `json:"emotionstatus"`
	Gender                   int      `json:"gender"`
	GenderString             string   `json:"genderstring"`
	HasIntroduce             bool     `json:"hasintroduce"`
	HeightString             string   `json:"heightstring"`
	HideVerifyModule         bool     `json:"hideverifymodule"`
	IntroduceContent         string   `json:"introducecontent"`
	IntroducePraiseCount     int      `json:"introducepraisecount"`
	IsActive                 bool     `json:"isactive"`
	IsFollowing              bool     `json:"isfollowing"`
	IsInBlackList            bool     `json:"isinblacklist"`
	IsStar                   bool     `json:"isstar"`
	IsZhenaiMail             bool     `json:"iszhenaimail"`
	LastLoginTimeString      string   `json:"lastlogintimestring"`
	LiveAudienceCount        int      `json:"liveaudiencecount"`
	LiveType                 int      `json:"livetype"`
	MarriageString           string   `json:"marriagestring"`
	MemberID                 int64    `json:"memberid"`
	MomentCount              int      `json:"momentcount"`
	Nickname                 string   `json:"nickname"`
	ObjectAgeString          string   `json:"objectagestring"`
	ObjectBodyString         string   `json:"objectbodystring"`
	ObjectChildrenString     string   `json:"objectchildrenstring"`
	ObjectEducationString    string   `json:"objecteducationstring"`
	ObjectHeightString       string   `json:"objectheightstring"`
	ObjectInfo               []string `json:"objectinfo"`
	ObjectMarriageString     string   `json:"objectmarriagestring"`
	ObjectSalaryString       string   `json:"objectsalarystring"`
	ObjectWantChildrenString string   `json:"objectwantchildrenstring"`
	ObjectWorkCityString     string   `json:"objectworkcitystring"`
	Onlive                   int      `json:"onlive"`
	PhotoCount               int      `json:"photocount"`
	Photos                   []Photos `json:"photos"`
	PraisedIntroduce         bool     `json:"praisedintroduce"`
	PreviewPhotoURL          string   `json:"previewphotourl"`
	SalaryString             string   `json:"salarystring"`
	ShowValidateIDCardFlag   bool     `json:"showvalidateidcardflag"`
	TotalPhotoCount          int      `json:"totalphotocount"`
	ValidateEducation        bool     `json:"validateeducation"`
	ValidateFace             bool     `json:"validateface"`
	ValidateIDCard           bool     `json:"validateidcard"`
	VideoCount               int      `json:"videocount"`
	VideoID                  int      `json:"videoid"`
	WorkCity                 int64    `json:"workcity"`
	WorkCityString           string   `json:"workcitystring"`
	WorkProvinceCityString   string   `json:"workprovincecitystring"`
}

type Photos struct {
	CreateTime  string `json:"createtime"`
	IsAvatar    bool   `json:"isavatar"`
	PhotoID     int64  `json:"photoid"`
	PhotoType   int    `json:"phototype"`
	PhotoURL    string `json:"photourl"`
	PraiseCount int    `json:"praisecount"`
	Praised     bool   `json:"praised"`
	Verified    bool   `json:"verified"`
}
