package command

type Droplet struct {
	Id           int      `json:"id"`
	Name         string      `json:"name"`
	Memory       int      `json:"memory"`
	Vcpus        int      `json:"vcpus"`
	Locked       bool      `json:"locked"`
	Status       string      `json:"status"`
	Kernel       struct {
			     Id      int      `json:"id"`
			     Name    string      `json:"name"`
			     Version string      `json:"version"`
		     } `json:"kernel"`
	created_at   string      `json:"created_at"`
	Backup_ids   []int      `json:"backup_ids"`
	Snapshot_ids []int      `json:"snapshot_ids"`
	Image        struct {
			     Id             int      `json:"id"`
			     Name           string      `json:"name"`
			     Distribution   string      `json:"distribution"`
			     Slug           string      `json:"slug"`
			     Public         bool      `json:"public"`
			     Regions        []string      `json:"regions"`
			     Created_at     string      `json:"created_at"`
			     Min_disk_size  int      `json:"min_disk_size"`
			     Itype          string      `json:"type"`
			     Size_gigabytes float32      `json:"size_gigabytes"`
		     }  `json:"image"`
}
type DropletsList struct {
	Id       string      `json:"id"`
	Droplets [] Droplet   `json:"droplets"`
}
