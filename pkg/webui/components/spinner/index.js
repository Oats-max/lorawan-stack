// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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
import bind from 'autobind-decorator'
import classnames from 'classnames'
import PropTypes from 'prop-types'

import from from '@ttn-lw/lib/from'

import style from './spinner.styl'

const id = () => `grad-${Math.round(Math.random() * 10000)}`

export default class Spinner extends React.PureComponent {
  static propTypes = {
    after: PropTypes.number,
    center: PropTypes.bool,
    children: PropTypes.node,
    className: PropTypes.string,
    faded: PropTypes.bool,
    inline: PropTypes.bool,
    micro: PropTypes.bool,
    small: PropTypes.bool,
  }

  static defaultProps = {
    after: 350,
    center: false,
    children: undefined,
    className: undefined,
    faded: false,
    inline: false,
    micro: false,
    small: false,
  }

  constructor(props) {
    super(props)

    this.state = { visible: false }
    this.id = id()
  }

  componentDidMount() {
    this.timer = setTimeout(this.show, this.props.after)
  }

  componentWillUnmount() {
    clearTimeout(this.timer)
  }

  @bind
  show() {
    this.setState({
      visible: true,
    })
  }

  render() {
    const {
      className,
      center,
      small,
      micro = false,
      faded = false,
      children,
      inline = false,
    } = this.props

    const { visible = false } = this.state

    const classname = classnames(
      style.box,
      className,
      ...from(style, {
        center,
        small,
        micro,
        faded,
        visible,
        inline,
      }),
    )

    return (
      <div className={classname}>
        <svg className={style.spinner} viewBox="0 0 100 100">
          <defs>
            <linearGradient id={this.id}>
              <stop offset="0%" className={style.stop} />
              <stop offset="100%" className={style.stop} stopColor="white" stopOpacity="0" />
            </linearGradient>
          </defs>
          <g transform="translate(50, 50)">
            <circle
              cx="0"
              cy="0"
              r={micro ? 35 : 40}
              className={style.bar}
              stroke={`url(#${this.id})`}
            />
          </g>
          <circle cx="50" cy="50" r={micro ? 35 : 40} className={style.circle} />
        </svg>
        <div className={style.message}>{children}</div>
      </div>
    )
  }
}
