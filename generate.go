package main

//go:generate go-bindata -pkg files_assets -o public/assets/files/files.go apps/files/css/... apps/files/js/... apps/files/img/...

//go:generate go-bindata -pkg federatedfilesharing_assets -o public/assets/federatedfilesharing/federatedfilesharing.go apps/federatedfilesharing/css/... apps/federatedfilesharing/js/... apps/federatedfilesharing/img/...

//go:generate go-bindata -pkg files_sharing_assets -o public/assets/files_sharing/files_sharing.go apps/files_sharing/css/... apps/files_sharing/js/... apps/files_sharing/img/...

//go:generate go-bindata -pkg files_trashbin_assets -o public/assets/files_trashbin/files_trashbin.go apps/files_trashbin/css/... apps/files_trashbin/img/... apps/files_trashbin/js/...

//go:generate go-bindata -pkg files_videoplayer_assets -o public/assets/files_videoplayer/files_videoplayer.go apps/files_videoplayer/...

//go:generate go-bindata -pkg core_assets -o public/assets/core/core.go core/css/... core/js/... core/img/... core/fonts/... core/search/css/... core/search/js/... core/search/templates/... core/vendor/...

//go:generate go-bindata -pkg gallery_assets -o public/assets/gallery/gallery.go apps/gallery/css/... apps/gallery/img/... apps/gallery/js/...

//go:generate go-bindata -pkg settings_assets -o public/assets/settings/settings.go settings/css/... settings/img/... settings/js/...

//go:generate go-bindata -pkg components -o public/assets/components/components.go public/components/...

//go:generate go-bindata -pkg templates_assets -o public/assets/templates/templates.go index.html
