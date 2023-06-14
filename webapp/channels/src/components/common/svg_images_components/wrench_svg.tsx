// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

type SvgProps = {
    width?: number;
    height?: number;
}

const SvgComponent = (props: SvgProps) => (
    <svg
        width={props.width ? props.width.toString() : '260'}
        height={props.height ? props.height.toString() : '260'}
        fill='var(--center-channel-bg)'
        viewBox='0 0 260 260'
        xmlns='http://www.w3.org/2000/svg'
    >
        <path
            d='M192.397 192.367L222.776 184.156L258.559 219.979C260.648 211.241 260.463 202.112 258.022 193.466C255.581 184.819 250.966 176.943 244.616 170.588C238.267 164.233 230.395 159.613 221.753 157.168C213.11 154.722 203.986 154.534 195.25 156.62L103.339 64.6975C105.407 55.9717 105.209 46.8616 102.765 38.2338C100.32 29.6059 95.71 21.7471 89.3726 15.4047C83.0352 9.06235 75.1814 4.44721 66.558 1.99823C57.9346 -0.450755 48.8282 -0.652137 40.105 1.41319L75.8885 37.2359L67.6308 67.5512L37.2523 75.7121L1.44389 39.8894C-0.648364 48.633 -0.464087 57.7673 1.97905 66.4193C4.42219 75.0713 9.04252 82.952 15.3986 89.3081C21.7546 95.6641 29.634 100.283 38.2836 102.724C46.9331 105.165 56.0638 105.345 64.8031 103.249L156.714 195.171C154.618 203.914 154.799 213.048 157.239 221.701C159.678 230.354 164.296 238.237 170.649 244.595C177.003 250.954 184.88 255.576 193.529 258.02C202.177 260.464 211.308 260.649 220.048 258.555L184.265 222.733L192.397 192.367Z'
            fill='#767D93'
        />
        <path
            d='M192.397 192.367L222.775 184.156L258.559 219.979C260.648 211.241 260.463 202.112 258.022 193.466C255.581 184.819 250.965 176.943 244.616 170.588C238.266 164.233 230.395 159.613 221.752 157.168C213.11 154.722 203.986 154.534 195.25 156.62L103.339 64.6975C105.407 55.9717 105.209 46.8616 102.764 38.2338C100.32 29.6059 95.7097 21.7471 89.3723 15.4047C83.0349 9.06235 75.1811 4.44721 66.5577 1.99823C57.9343 -0.450755 48.8279 -0.652137 40.1047 1.41319L75.8882 37.2359L67.6305 67.5512'
            fill='#A4A9B7'
        />
    </svg>
);

export default SvgComponent;
