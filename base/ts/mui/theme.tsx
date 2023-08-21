import React from 'react';
import { createTheme, ThemeOptions } from '@mui/material/styles';
import { Link as RouterLink, LinkProps as RouterLinkProps } from 'react-router-dom';
import { LinkProps } from '@mui/material/Link';

const LinkBehavior = React.forwardRef<
  HTMLAnchorElement,
  Omit<RouterLinkProps, 'to'> & { href: RouterLinkProps['to'] }
>((props, ref) => {
  const { href, ...other } = props;
  // If the element indicates "real" link, use <a> tag
  // Determine based on className
  if (other.className?.includes('native-link')) {
    return <a ref={ref} href={href} {...other} />;
  }
  // Map href (Material UI) -> to (react-router)
  return <RouterLink ref={ref} to={href} {...other} />;
});

export const peridotThemeOptions: ThemeOptions = {
  palette: {
    mode: 'light',
    primary: {
      main: '#1c2834',
    },
    secondary: {
      main: '#168afd',
    },
  },
  components: {
    MuiLink: {
      defaultProps: {
        component: LinkBehavior,
      } as LinkProps,
    },
    MuiButtonBase: {
      defaultProps: {
        LinkComponent: LinkBehavior,
      },
    },
  },
};

export const peridotTheme = createTheme(peridotThemeOptions);
