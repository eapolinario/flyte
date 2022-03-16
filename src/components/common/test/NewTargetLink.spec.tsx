import { render } from '@testing-library/react';
import * as React from 'react';
import { NewTargetLink } from '../NewTargetLink';

const linkTarget = 'https://www.github.com/lyft/flyteconsole';

const linkText = 'open in new page';
const renderLink = () => render(<NewTargetLink href={linkTarget}>{linkText}</NewTargetLink>);
const renderExternalLink = () =>
  render(
    <NewTargetLink href={linkTarget} external={true}>
      {linkText}
    </NewTargetLink>,
  );

test('renders a blank target link', () => {
  const { getByText } = renderLink();
  const anchor = getByText(linkText);
  expect(anchor).toHaveAttribute('target', '_blank');
});

test('renders with additional icon for external links', () => {
  const { container } = renderExternalLink();
  expect(container.querySelector('svg')).not.toBeNull();
});
