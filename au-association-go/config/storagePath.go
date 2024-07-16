package config

type StoragePath struct {
	RootActivityPath   string `mapstructure:"RootActivityPath" json:"RootActivityPath" yaml:"RootActivityPath"`
	RootAssLogoPath    string `mapstructure:"RootAssLogoPath" json:"RootAssLogoPath" yaml:"RootAssLogoPath"`
	RootOutlayPath     string `mapstructure:"RootOutlayPath" json:"RootOutlayPath" yaml:"RootOutlayPath"`
	RootPicturePath    string `mapstructure:"RootPicturePath" json:"RootPicturePath" yaml:"RootPicturePath"`
	RootSynthesizePath string `mapstructure:"RootSynthesizePath" json:"RootSynthesizePath" yaml:"RootSynthesizePath"`
	RootImagesPath     string `mapstructure:"RootImagesPath" json:"RootImagesPath" yaml:"RootImagesPath"`
}
