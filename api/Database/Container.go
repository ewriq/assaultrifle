package Database

import (
	"assaultrifle/Container"
	"assaultrifle/Form"
	"fmt"
)

/*
type name password port token (container id) user

listmybase
listALL



*/

func ContainerAdd(name, password, port, user, types string) (bool, string) {
	if !FinderPort(port) {
		if types == "MySql" {
			pull := Container.PullImage("mysql:8.0")
			if pull == nil {
			Token, errr := Container.CreateMySQLContainer(name, "mysql:8.0", port, password)
			if errr == nil {
				query := "INSERT INTO container (name, password, port, user, types, token) VALUES (?, ?, ?, ?, ?,?)"
				_, err := db.Exec(query, name, password, port, user, "MySql", Token)
				if err != nil {
					fmt.Println(err)
					return false, ""
				}
				return true, Token
			}
			}

		}

		if types == "PostgreSql" {
			pull := Container.PullImage("postgres:16")
			if pull == nil {
			Token, errr := Container.CreatePostgresContainer(name, "postgres:16", port, password)
			if errr == nil {
				query := "INSERT INTO container (name, password, port, user, types, token) VALUES (?, ?, ?, ?, ?,?)"
				_, err := db.Exec(query, name, password, port, user, "PostgreSql", Token)
				if err != nil {
					fmt.Println(err)
					return false, ""
				}
				return true, Token
			}
			}
		}

		if types == "NodeJS" {
			pull := Container.PullImage("node:18")
			if pull == nil {
				Token, errr := Container.CreateNodeContainer(name, "node:18", port)
				if errr == nil {
					query := "INSERT INTO container (name, password, port, user, types, token) VALUES (?, ?, ?, ?, ?,?)"
					_, err := db.Exec(query, name, password, port, user, "Nextcloud", Token)
					if err != nil {
						fmt.Println(err)
						return false, ""
					}
					return true, Token
				}
			}
		}

		
	}
	return false, ""
}

func FinderPort(port string) bool {
	query := "SELECT COUNT(*) FROM container WHERE port = ?"
	row := db.QueryRow(query, port)

	var detected int
	err := row.Scan(&detected)
	if err != nil {
		return false
	}

	if detected == 0 {
		return false
	}

	return true
}

func ContainerDelete(token string) (bool, string) {
	err := Container.DeleteContainer(token)
	if err != nil {
		fmt.Println("❌ Docker silme hatası:", err)
		return false, ""
	}
	query := "DELETE FROM container WHERE token = ?"
	_, err = db.Exec(query, token)
	if err != nil {
		fmt.Println("❌ DB silme hatası:", err)
		return false, ""
	}

	return true, token
}

func ContainerListMy(user string) ([]Form.Container, error) {
	query := "SELECT name, password, port, token, user, types, id FROM container WHERE user = ?"
	rows, err := db.Query(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var containers []Form.Container

	for rows.Next() {
		var container Form.Container
		err := rows.Scan(
			&container.Name,
			&container.Password,
			&container.Port,
			&container.Token,
			&container.User,
			&container.Type,
			&container.ID,
		)
		if err != nil {
			return nil, err
		}
		containers = append(containers, container)
	}

	if len(containers) == 0 {
		return nil, fmt.Errorf("hiç container bulunamadı")
	}

	return containers, nil
}
func ContainerListAll(user string) ([]Form.Container, error) {
	query := "SELECT name, password, port, token, user, types, id FROM container WHERE user = ?"
	rows, err := db.Query(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var containers []Form.Container

	for rows.Next() {
		var container Form.Container
		err := rows.Scan(
			&container.Name,
			&container.Password,
			&container.Port,
			&container.Token,
			&container.User,
			&container.Type,
			&container.ID,
		)
		if err != nil {
			return nil, err
		}
		containers = append(containers, container)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(containers) == 0 {
		return nil, fmt.Errorf("container bulunamadı")
	}

	return containers, nil
}
