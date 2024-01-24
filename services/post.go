package services

import (
	"errors"

	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/models"
)

type PostForm struct {
	Title	   string `json:"title"`
	Body	   string `json:"body"`
}

type PostService struct{}

func (srv PostService) CreatePost(userID uint, postForm PostForm) error {
	// Create post
	result := initializers.DB.Create(&models.Posts{
		Title:    postForm.Title,
		Body:     postForm.Body,
		User_id:  userID,
	})

	if result.Error != nil {
		err := errors.New("Failed to create post")
		return err
	}

	return nil
}

func (srv PostService) GetUserPosts(userID uint) ([]models.Posts, error) {
	var posts []models.Posts

	result := initializers.DB.Where("user_id =?", userID).Find(&posts)

	if result.Error != nil {
		err := errors.New("Failed to get all posts")
		return nil, err
	}

	if result.RowsAffected == 0 {
		err := errors.New("There are no posts")
		return nil, err
	}

	return posts, nil
}

func (srv PostService) GetOnePost(postID uint) (models.Posts, error) {
	var post models.Posts

	result := initializers.DB.First(&post, postID)
	if result.Error != nil {
		err := errors.New("Failed to get post")
		return post, err
	}

	if result.RowsAffected == 0 {
		err := errors.New("There are no post")
		return post, err
	}

	return post, nil
}

func (srv PostService) UpdatePost(userID uint, postID uint, postForm PostForm) error{
	//Check whether post is
	var post models.Posts
	result1 := initializers.DB.First(&post, postID)
	if result1.Error != nil{
		err := errors.New("There are no post that match postID")
		return err
	}

	//Check whether post is user's post
	if post.User_id != userID {
		err := errors.New("You are not this post's creater")
		return err
	}

	//Update post
	post.Title = postForm.Title
	post.Body = postForm.Body
	result2 := initializers.DB.Save(&post)

	if result2.Error != nil {
		err := errors.New("Failed to update post")
		return err
	}

	return nil
}