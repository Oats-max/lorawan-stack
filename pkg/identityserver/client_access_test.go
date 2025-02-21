// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identityserver

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/storetest"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"google.golang.org/grpc"
)

func TestClientCollaborators(t *testing.T) {
	p := &storetest.Population{}

	admin := p.NewUser()
	admin.Admin = true
	adminKey, _ := p.NewAPIKey(admin.GetEntityIdentifiers(), ttnpb.Right_RIGHT_ALL)
	adminCreds := rpcCreds(adminKey)

	usr1 := p.NewUser()
	usr1Key, _ := p.NewAPIKey(usr1.GetEntityIdentifiers(), ttnpb.Right_RIGHT_ALL)
	usr1Creds := rpcCreds(usr1Key)

	cli1 := p.NewClient(usr1.GetOrganizationOrUserIdentifiers())

	usr2 := p.NewUser()
	p.NewMembership(
		usr2.GetOrganizationOrUserIdentifiers(),
		cli1.GetEntityIdentifiers(),
		ttnpb.Right_RIGHT_CLIENT_ALL,
	)

	usr3 := p.NewUser()

	t.Parallel()
	a, ctx := test.New(t)

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		is.config.AdminRights.All = true

		reg := ttnpb.NewClientAccessClient(cc)

		// GetCollaborator that doesn't exist.
		got, err := reg.GetCollaborator(ctx, &ttnpb.GetClientCollaboratorRequest{
			ClientIds:    cli1.GetIds(),
			Collaborator: usr3.GetOrganizationOrUserIdentifiers(),
		}, adminCreds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
		a.So(got, should.BeNil)

		// SetCollaborator without credentials.
		_, err = reg.SetCollaborator(ctx, &ttnpb.SetClientCollaboratorRequest{
			ClientIds: cli1.GetIds(),
			Collaborator: &ttnpb.Collaborator{
				Ids: usr2.GetOrganizationOrUserIdentifiers(),
				Rights: []ttnpb.Right{
					ttnpb.Right_RIGHT_CLIENT_ALL,
				},
			},
		})
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		// GetCollaborator without credentials.
		got, err = reg.GetCollaborator(ctx, &ttnpb.GetClientCollaboratorRequest{
			ClientIds:    cli1.GetIds(),
			Collaborator: usr2.GetOrganizationOrUserIdentifiers(),
		})
		if a.So(err, should.NotBeNil) && a.So(errors.IsPermissionDenied(err), should.BeTrue) {
			a.So(got, should.BeNil)
		}

		// ListCollaborators without credentials.
		list, err := reg.ListCollaborators(ctx, &ttnpb.ListClientCollaboratorsRequest{
			ClientIds: cli1.GetIds(),
		})
		if a.So(err, should.NotBeNil) && a.So(errors.IsUnauthenticated(err), should.BeTrue) {
			a.So(list, should.BeNil)
		}

		// Collaborator CRUD with different valid credentials.
		for _, opts := range [][]grpc.CallOption{{adminCreds}, {usr1Creds}} {
			_, err := reg.SetCollaborator(ctx, &ttnpb.SetClientCollaboratorRequest{
				ClientIds: cli1.GetIds(),
				Collaborator: &ttnpb.Collaborator{
					Ids: usr3.GetOrganizationOrUserIdentifiers(),
					Rights: []ttnpb.Right{
						ttnpb.Right_RIGHT_CLIENT_ALL,
					},
				},
			}, opts...)
			a.So(err, should.BeNil)

			list, err := reg.ListCollaborators(ctx, &ttnpb.ListClientCollaboratorsRequest{
				ClientIds: cli1.GetIds(),
			}, opts...)
			if a.So(err, should.BeNil) && a.So(list, should.NotBeNil) && a.So(list.Collaborators, should.HaveLength, 3) {
				for _, k := range list.Collaborators {
					if unique.ID(ctx, k.GetIds()) == unique.ID(ctx, usr3.GetIds()) {
						a.So(k.Rights, should.Resemble, []ttnpb.Right{
							ttnpb.Right_RIGHT_CLIENT_ALL,
						})
					}
				}
			}

			got, err := reg.GetCollaborator(ctx, &ttnpb.GetClientCollaboratorRequest{
				ClientIds:    cli1.GetIds(),
				Collaborator: usr3.GetOrganizationOrUserIdentifiers(),
			}, opts...)
			if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
				a.So(got.Rights, should.Resemble, []ttnpb.Right{
					ttnpb.Right_RIGHT_CLIENT_ALL,
				})
			}

			_, err = reg.SetCollaborator(ctx, &ttnpb.SetClientCollaboratorRequest{
				ClientIds: cli1.GetIds(),
				Collaborator: &ttnpb.Collaborator{
					Ids: usr3.GetOrganizationOrUserIdentifiers(),
				},
			}, opts...)
			a.So(err, should.BeNil)

			got, err = reg.GetCollaborator(ctx, &ttnpb.GetClientCollaboratorRequest{
				ClientIds:    cli1.GetIds(),
				Collaborator: usr3.GetOrganizationOrUserIdentifiers(),
			}, opts...)
			if a.So(err, should.NotBeNil) {
				a.So(errors.IsNotFound(err), should.BeTrue)
			}
			a.So(got, should.BeNil)
		}

		// Remove the other collaborator with _ALL rights.
		_, err = reg.SetCollaborator(ctx, &ttnpb.SetClientCollaboratorRequest{
			ClientIds: cli1.GetIds(),
			Collaborator: &ttnpb.Collaborator{
				Ids: usr2.GetOrganizationOrUserIdentifiers(),
			},
		}, usr1Creds)
		a.So(err, should.BeNil)

		// Try removing the only collaborator with _ALL rights.
		_, err = reg.SetCollaborator(ctx, &ttnpb.SetClientCollaboratorRequest{
			ClientIds: cli1.GetIds(),
			Collaborator: &ttnpb.Collaborator{
				Ids: usr1.GetOrganizationOrUserIdentifiers(),
			},
		}, usr1Creds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsFailedPrecondition(err), should.BeTrue)
		}
	}, withPrivateTestDatabase(p))
}

func TestClientAccessClusterAuth(t *testing.T) {
	p := &storetest.Population{}
	cli1 := p.NewClient(nil)

	t.Parallel()
	a, ctx := test.New(t)

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttnpb.NewClientAccessClient(cc)

		rights, err := reg.ListRights(ctx, cli1.GetIds(), is.WithClusterAuth())
		if a.So(err, should.BeNil) && a.So(rights, should.NotBeNil) {
			a.So(ttnpb.AllClusterRights.Intersect(ttnpb.AllClientRights).Sub(rights).Rights, should.BeEmpty)
		}
	}, withPrivateTestDatabase(p))
}
