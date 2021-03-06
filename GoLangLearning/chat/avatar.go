package main

import (
	"errors"
	"path"
	"io/ioutil"
/*	"strings"
	"fmt"
/*	"crypto/md5"*/
)

var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

// Avatar represents types capable of representing
// user profile pictures.
type Avatar interface {

	//GetAvatarURL(ChatUser) (string, error)
	GetAvatarURL(c *client) (string, error)
}

/*
type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}*/

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (_ FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	
	if userid, ok := c.userData["userid"]; ok{
		if useridStr, ok := userid.(string); ok{
			//return "/avatars/" + useridStr + ".jpg", nil
			if files, err := ioutil.ReadDir("avatars"); err == nil{
				for _, file := range files {
					if file.IsDir(){
						continue
					}
					if match, _ := path.Match(useridStr + "*", file.Name()); match{
						return "/avatars/" + file.Name(), nil
					}
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	/*url := u.AvatarURL()
	if len(url) > 0 {
		return u.AvatarURL(), nil
	}
	return "", ErrNoAvatarURL*/
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr,  nil
		}
	}
	return "", ErrNoAvatarURL
}
