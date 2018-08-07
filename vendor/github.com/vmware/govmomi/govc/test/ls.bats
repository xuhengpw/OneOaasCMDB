#!/usr/bin/env bats

load test_helper

@test "ls" {
  run govc ls
  assert_success
  # /dc/{vm,network,host,datastore}
  [ ${#lines[@]} -ge 4 ]

  run govc ls host
  assert_success
  [ ${#lines[@]} -ge 1 ]

  run govc ls enoent
  assert_success
  [ ${#lines[@]} -eq 0 ]
}

@test "ls vm" {
  vm=$(new_empty_vm)

  run govc ls vm
  assert_success
  [ ${#lines[@]} -ge 1 ]

  run govc ls vm/$vm
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run govc ls /*/vm/$vm
  assert_success
  [ ${#lines[@]} -eq 1 ]
}

@test "ls network" {
  run govc ls network
  assert_success
  [ ${#lines[@]} -ge 1 ]

  local path=${lines[0]}
  run govc ls "$path"
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run govc ls "network/$(basename "$path")"
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run govc ls "/*/network/$(basename "$path")"
  assert_success
  [ ${#lines[@]} -eq 1 ]
}

@test "ls multi ds" {
  vcsim_env

  run govc ls
  assert_success
  # /DC0/{vm,network,host,datastore}
  [ ${#lines[@]} -eq 4 ]

  run govc ls /DC*
  assert_success
  # /DC[0,1]/{vm,network,host,datastore}
  [ ${#lines[@]} -eq 8 ]

  # here 'vm' is relative to /DC0
  run govc ls vm
  assert_success
  [ ${#lines[@]} -gt 0 ]

  unset GOVC_DATACENTER

  run govc ls
  assert_success
  # /DC[0,1]
  [ ${#lines[@]} -eq 2 ]

  run govc ls -dc enoent
  assert_failure
  [ ${#lines[@]} -gt 0 ]

  # here 'vm' is relative to '/' - so there are no matches
  run govc ls vm
  assert_success
  [ ${#lines[@]} -eq 0 ]

  # ls all vms in all datacenters
  run govc ls */vm
  assert_success
  [ ${#lines[@]} -gt 0 ]
}

@test "ls moref" {
    # ensure the vm folder isn't empty
    run govc vm.create -on=false "$(new_id)"
    assert_success

    # list dc folder paths
    folders1=$(govc ls)
    # list dc folder refs | govc ls -L ; should output the same paths
    folders2=$(govc ls -i | xargs govc ls -L)

    assert_equal "$folders1" "$folders2"

    for folder in $folders1
    do
        # list paths in $folder
        items1=$(govc ls "$folder")
        # list refs in $folder | govc ls -L ; should output the same paths
        items2=$(govc ls -i "$folder" | xargs -d '\n' govc ls -L)

        assert_equal "$items1" "$items2"
    done
}
