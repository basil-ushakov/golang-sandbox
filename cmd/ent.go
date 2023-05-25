package main

import (
  "context"
  "fmt"
  "database/sql"
  _ "os"

  "log"

  "time"

  "github.com/NICKNAME-wengreen/BigDemo/ent"
  "github.com/NICKNAME-wengreen/BigDemo/ent/user"
  "github.com/NICKNAME-wengreen/BigDemo/ent/car"
  "github.com/NICKNAME-wengreen/BigDemo/ent/group"

  "entgo.io/ent/dialect"
  entsql "entgo.io/ent/dialect/sql"
  _ "github.com/jackc/pgx/v5/stdlib"
)

type ServerHandler interface {
  db_connect(ctx context.Context, databaseUrl string) *ent.Client
}

type ServerType struct {
  client *ent.Client
}

func (sh *ServerType) db_connect(ctx context.Context,databaseUrl string) {
  dbpool, err := sql.Open("pgx", databaseUrl)
  if err != nil {
    log.Fatal(err)
  }
  drv    := entsql.OpenDB(dialect.Postgres, dbpool)
  client := ent.NewClient(ent.Driver(drv))
  sh.client = client

}

func crt_user(ctx context.Context, client *ent.Client,name string) (*ent.User, error) {
  user,err := client.User.
    Create().
    SetAge(30).
    SetName(name).
    Save(ctx)

  if err != nil {
    return nil, fmt.Errorf("failed creating user: %w", err)
  }

  return user, nil
}

func crt_cars(ctx context.Context, client *ent.Client) error {
	var car_names = [4]string{
		"CyberTruck"	,
		"Ford"		,
		"Nissan Skyline",
		"Porche 911"	,
	}

	for _, car_name := range car_names {
		vehicle, err := client.Car.
			Create().
			SetModel(car_name).
			SetRegisteredAt(time.Now()).
			// SetUserCars("a8m").
			Save(ctx)

		if err != nil {
			return fmt.Errorf("failed creating car: %w", err)
		}
		log.Println("car was created: ", vehicle)
	}

	return nil
}

type User_T struct {
	age int
	name string
}

func crt_graph(ctx context.Context, client *ent.Client) error {
	userData_slice := []User_T{
		User_T{
			age: 20,
			name: "E6|lan",
		},
		User_T{
			age: 25,
			name: "Ye6ak",
		},
		User_T{
			age: 30,
			name: "Kycok_7oBHa",
		},
		User_T{
			age: 43,
			name: "Mood_Ak",
		},
	}

	user_slice := []*ent.User{}

	for _, item := range userData_slice {
		user, err := client.User.
			Create().
			SetAge(item.age).
			SetName(item.name).
			Save(ctx)

			user_slice = append(user_slice,user)
		if err != nil {
			return err
		}
	}

	carData_slice := []string{
		"Mazda",
		"Tesla",
		"Mustang",
		"Viper",
	}

	for ind, item := range carData_slice {
		err := client.Car.
			Create().
			SetModel(item).
			SetRegisteredAt(time.Now()).
			SetOwner(user_slice[ind]).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	var err error

	err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(user_slice[0],user_slice[1]).
		Exec(ctx)
	if err != nil {
		return err
	}


	err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(user_slice[2],user_slice[3]).
		Exec(ctx)
	if err != nil {
		return err
	}

	log.Println("The graph was created successfully")

	return nil
}


func qry_user(ctx context.Context, client *ent.Client,name string) (*ent.User, error) {
  user, err := client.User.
	Query().
	Where(user.Name(name)).
	Only(ctx)

  if err != nil {
    return nil, fmt.Errorf("failed querying user: %w", err)
  }
  log.Println("user returned: ",user)
  return user, nil
}

func qry_carsUser(ctx context.Context, user *ent.User) error {
	cars, err := user.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}

	for _, c := range cars {
		owner, err := c.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", c.Model, err)
		}
		log.Printf("car %q owner: %q\n", c.Model, owner.Name)
	}

	return nil
}

func qry_groupName(ctx context.Context, client *ent.Client,name string) error {
	cars, err := client.Group.
		Query().
		Where(group.Name(name)).
		QueryUsers().
		QueryCars().
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)

	return nil
}

func qry_carsUsersByName(ctx context.Context, client *ent.Client,name string) error {
	user := client.User.
		Query().
		Where(
			user.HasCars(),
			user.Name(name),
		).
		OnlyX(ctx)

	cars, err := user.
		QueryGroups().
		QueryUsers().
		QueryCars().
		Where(
			car.Not(
				car.Model("Mazda"),
			),
		).
		All(ctx)

	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	return nil
}

func qry_groupByUsers(ctx context.Context,client *ent.Client) error {
	groups, err := client.Group.
		Query().
		Where(group.HasUsers()).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting groups: %w", err)
	}
	log.Println("groups returned:", groups)
	return nil
}

func main() {
  ctx    := context.Background()

  var servhd ServerType
      servhd.db_connect(ctx,"postgres://bitterman:sosiska228@localhost:6666/quicker")

  client := servhd.client
  if err := client.Schema.Create(ctx); err != nil {
    log.Fatal(err)
  }

  // var crnt_user *ent.User
  var err error

  // var name string = "E6|lan"

  // crt_user(ctx,client,name)
  // crt_cars(ctx,client)
  //
  // if err = crt_graph(ctx,client); err != nil {
  //   log.Fatal(err)
  // }
  //
  // if crnt_user,err = qry_user(ctx,client,name); err != nil {
  //   log.Fatal(err)
  // }
  //
  // if err = qry_carsUser(ctx,crnt_user); err != nil {
  //   log.Fatal(err)
  // }
  //
  // if err = qry_groupName(ctx,client,"GitLab"); err != nil {
  //   log.Fatal(err)
  // }
  //
  // if err = qry_carsUsersByName(ctx,client,name); err != nil {
  //   log.Fatal(err)
  // }

  if err = qry_groupByUsers(ctx,client); err != nil {
    log.Fatal(err)
  }

  log.Println("DB connection established")
  defer client.Close()
}
