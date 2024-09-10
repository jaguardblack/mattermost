// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {render} from '@testing-library/react';
import React from 'react';
import {act} from 'react-dom/test-utils';

import LatexInline from 'components/latex_inline/latex_inline';

import {withIntl} from 'tests/helpers/intl-test-helper';

describe('components/LatexInline', () => {
    const defaultProps = {
        content: 'e^{i\\pi} + 1 = 0',
        enableInlineLatex: true,
    };

    test('should match snapshot', async () => {
        let container;

        await act(async () => {
            const result = render(withIntl(<LatexInline {...defaultProps}/>));
            container = result.container;
        });
        expect(container).toMatchSnapshot();
    });

    test('latex is disabled', async () => {
        const props = {
            ...defaultProps,
            enableInlineLatex: false,
        };

        let container;

        await act(async () => {
            const result = render(withIntl(<LatexInline {...props}/>));
            container = result.container;
        });
        expect(container).toMatchSnapshot();
    });

    test('error in katex', async () => {
        const props = {
            content: 'e^{i\\pi + 1 = 0',
            enableInlineLatex: true,
        };

        let container;

        await act(async () => {
            const result = render(withIntl(<LatexInline {...props}/>));
            container = result.container;
        });
        expect(container).toMatchSnapshot();
    });
});
