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

import React from 'react'
import { Switch, Route } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ErrorView from '@ttn-lw/lib/components/error-view'
import NotFoundRoute from '@ttn-lw/lib/components/not-found-route'

import SubViewError from '@console/views/sub-view-error'
import OrganizationCollaboratorsList from '@console/views/organization-collaborators-list'
import OrganizationCollaboratorAdd from '@console/views/organization-collaborator-add'
import OrganizationCollaboratorEdit from '@console/views/organization-collaborator-edit'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'
import { userPathId as userPathIdRegexp } from '@ttn-lw/lib/regexp'

const OrganizationCollaborators = props => {
  const { orgId, match } = props

  useBreadcrumbs(
    'orgs.single.collaborators',
    <Breadcrumb
      path={`/organizations/${orgId}/collaborators`}
      content={sharedMessages.collaborators}
    />,
  )

  return (
    <ErrorView errorRender={SubViewError}>
      <Switch>
        <Route exact path={`${match.path}`} component={OrganizationCollaboratorsList} />
        <Route exact path={`${match.path}/add`} component={OrganizationCollaboratorAdd} />
        <Route
          path={`${match.path}/:collaboratorType(user|organization)/:collaboratorId${userPathIdRegexp}`}
          component={OrganizationCollaboratorEdit}
        />
        <NotFoundRoute />
      </Switch>
    </ErrorView>
  )
}
OrganizationCollaborators.propTypes = {
  match: PropTypes.match.isRequired,
  orgId: PropTypes.string.isRequired,
}

export default OrganizationCollaborators
