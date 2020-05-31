package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/entity"
)

const (
	userDetailByEmailKey = "get_user:%s"
	userDetailByEmailTtl = 3600
)

func (p *UserRepo) GetByUserInfoByEmail(ctx context.Context, user entity.User) (*entity.User, error) {
	stmt, err := p.DB.PrepareContext(ctx, getUserInfoByEmail)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, &user.Email)

	res := entity.User{}
	err = row.Scan(
		&res.ID,
		&res.Email,
		&res.FullName,
		&res.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (p *UserRepo) InsertUserInfo(ctx context.Context, user entity.User) (*entity.User, error) {
	stmt, err := p.DB.PrepareContext(ctx, insertUserInfo)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, &user.ID, &user.Email, &user.FullName, &user.PhoneNumber)

	res := entity.User{}
	err = row.Scan(
		&res.ID,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (p *UserRepo) UpdatetUserInfo(ctx context.Context, user entity.User) (*entity.User, error) {
	stmt, err := p.DB.PrepareContext(ctx, updateUserInfo)
	if err != nil {
		return nil, err
	}

	//row := stmt.QueryRowContext(ctx, &user.Email, &user.FullName, &user.PhoneNumber, &user.ID)
	result, err2 := stmt.Exec(&user.Email, &user.FullName, &user.PhoneNumber, &user.ID)

	if err2 != nil {
		return nil, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("THere are no data affected")
	}

	return &user, nil
}

func (p *UserRepo) DeleteUserInfo(ctx context.Context, user entity.User) (*entity.User, error) {
	stmt, err := p.DB.PrepareContext(ctx, deleteUserInfo)
	if err != nil {
		return nil, err
	}

	//row := stmt.QueryRowContext(ctx, &user.ID)
	result, err2 := stmt.Exec(&user.ID)

	if err2 != nil {
		return nil, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("THere are no data affected")
	}

	return &user, nil
}

func (p *UserRepo) SetUserDetailByEmailCache(ctx context.Context, user entity.User) error {
	key := fmt.Sprintf(userDetailByEmailKey, user.Email)
	marshall, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	c := p.Cache.Get()
	defer func() {
		err := c.Close()
		if nil != err {
			fmt.Println("Error redis : ", err)
		}
	}()
	_, err = c.Do("SET", key, marshall)
	if err != nil {
		return err
	}
	_, err = c.Do("EXPIRE", key, userDetailByEmailTtl)
	if err != nil {
		return err
	}
	return nil
}

func (p *UserRepo) GetUserDetailByEmailCache(ctx context.Context, user entity.User) (*entity.User, error) {
	key := fmt.Sprintf(userDetailByEmailKey, user.Email)
	c := p.Cache.Get()
	defer func() {
		err := c.Close()
		if nil != err {
			fmt.Println("Error redis : ", err)
		}
	}()
	getUser, err := c.Do("GET", key)
	if err != nil {
		return nil, err
	}
	// TODO: use scan struct
	//// Fetch all album fields with the HGETALL command. Wrapping this
	//// in the redis.Values() function transforms the response into type
	//// []interface{}, which is the format we need to pass to
	//// redis.ScanStruct() in the next step.
	//values, err := redis.Values(conn.Do("HGETALL", "album:1"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Create an instance of an Album struct and use redis.ScanStruct()
	//// to automatically unpack the data to the struct fields. This uses
	//// the struct tags to determine which data is mapped to which
	//// struct fields.
	//var album Album
	//err = redis.ScanStruct(values, &album)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//if empty then return
	if nil == getUser {
		return nil, nil
	}
	var resUser entity.User
	b, ok := getUser.([]uint8)
	if ok {
		err = json.Unmarshal(b, &resUser)
		if err != nil {
			return nil, err
		}
	}
	return &resUser, nil
}
