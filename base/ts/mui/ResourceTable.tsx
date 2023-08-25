import React from 'react';
import { Link, useSearchParams } from 'react-router-dom';

import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Table from '@mui/material/Table';
import Paper from '@mui/material/Paper';
import TableRow from '@mui/material/TableRow';
import TableBody from '@mui/material/TableBody';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import LinearProgress from '@mui/material/LinearProgress';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';

import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import ArrowForwardIcon from '@mui/icons-material/ArrowForward';
import InputLabel from '@mui/material/InputLabel';

import { StandardResource } from '../resource';

export interface ResourceTableField {
  key: string;
  label: string;
}

export interface ResourceTableProps<T> {
  fields: ResourceTableField[];

  // load is usually the OpenAPI SDK function that loads the resource.
  load(pageSize: number, pageToken?: string): Promise<any>;

  // transform can be used to transform the response from the load function.
  // usually for List functions, the response is usually wrapped in a
  // ListResponse object, so this function can be used to extract the list
  // from the response.
  transform?(response: { [key: string]: T }): T[];
}

export function ResourceTable<T extends StandardResource>(
  props: ResourceTableProps<T>,
) {
  const allowedPageSizes = [1, 25, 50, 100, 500, 1000];

  // State for pagination
  // We can use query parameters to store the page token, as well as the
  // page size and history of page tokens.
  const [search, setSearch] = useSearchParams();
  const pt = search.get('pt') || undefined;

  // The page token history is base64 encoded, so we need to decode it
  // (if it exists of course) and then parse it as JSON.
  // It should be an array of strings.
  const pth = search.get('pth');
  let initPageTokenHistory: string[] = [];
  if (pth) {
    initPageTokenHistory = JSON.parse(atob(pth));
  }

  let initRowsPerPage = parseInt(search.get('rpp') || '25') || 25;
  if (!allowedPageSizes.includes(initRowsPerPage)) {
    initRowsPerPage = 25;
  }

  const [pageToken, setPageToken] = React.useState<string | undefined>(pt);
  const [nextPageToken, setNextPageToken] = React.useState<string | undefined>(
    undefined,
  );
  const [pageTokenHistory, setPageTokenHistory] =
    React.useState<string[]>(initPageTokenHistory);
  const [rowsPerPage, setRowsPerPage] = React.useState<number>(initRowsPerPage);
  const [rows, setRows] = React.useState<T[] | undefined>(undefined);
  const [loading, setLoading] = React.useState<boolean>(false);

  // Update the query parameters when any of the pagination state changes
  React.useEffect(() => {
    const search = new URLSearchParams(location.search);
    if (pageToken) {
      search.set('pt', pageToken);
    } else {
      search.delete('pt');
    }
    search.set('rpp', rowsPerPage.toString());
    if (pageTokenHistory.length > 0) {
      search.set('pth', btoa(JSON.stringify(pageTokenHistory)));
    } else {
      search.delete('pth');
    }
    setSearch(search);
  }, [pageToken, pageTokenHistory]);

  // Rows per page changing means we need to reset the page token history
  React.useEffect(() => {
    const search = new URLSearchParams(location.search);
    search.set('rpp', rowsPerPage.toString());
    setSearch(search);
    setPageTokenHistory([]);
    setPageToken(undefined);
  }, [rowsPerPage]);

  // Load the resource using useEffect
  React.useEffect(() => {
    (async () => {
      setLoading(true);
      const [res, err] = await props.load(rowsPerPage, pageToken);
      setLoading(false);
      if (err) {
        console.log(err);
        return;
      }

      setNextPageToken(res.nextPageToken);

      if (props.transform) {
        setRows(props.transform(res));
      } else {
        setRows(res);
      }
    })().then();
  }, [pageToken, rowsPerPage]);

  // Create table header
  const header = props.fields.map((field) => {
    return <TableCell key={field.key}>{field.label}</TableCell>;
  });

  // Create table body
  const body = rows?.map((row) => (
    <TableRow
      key={row.name}
      sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
    >
      {props.fields.map((field) => {
        return (
          <TableCell key={field.key}>
            {field.key === 'name' ? (
              <Link to={`/${row.name}`}>{row.name}</Link>
            ) : (
              <>{row[field.key] ? row[field.key].toString() : '--'}</>
            )}
          </TableCell>
        );
      })}
    </TableRow>
  ));

  // Should show previous button if history is not empty
  let showPrevious = false;
  if (pageToken) {
    showPrevious = true;
  }

  // Should show next button if next page token is not empty
  let showNext = false;
  if (nextPageToken) {
    showNext = true;
  }

  // If loading disable both buttons
  if (loading) {
    showPrevious = false;
    showNext = false;
  }

  // Handle previous page button
  const handlePreviousPage = () => {
    if (pageTokenHistory.length > 0) {
      const newPageTokenHistory = [...pageTokenHistory];
      const prevPageToken = newPageTokenHistory.pop();
      setPageTokenHistory(newPageTokenHistory);
      setPageToken(prevPageToken);
    }

    // If the history is empty, then we this probably means this is page 2
    // and we should clear the page token.
    if (pageTokenHistory.length === 0) {
      setPageToken(undefined);
    }
  };

  // Handle next page button
  const handleNextPage = () => {
    if (nextPageToken) {
      if (pageToken) {
        setPageTokenHistory([...pageTokenHistory, pageToken]);
      }
      setPageToken(nextPageToken);
    }
  };

  return (
    <>
      <TableContainer component={Paper} elevation={2}>
        {loading && <LinearProgress />}
        <Table>
          <TableHead>
            <TableRow>{header}</TableRow>
          </TableHead>
          {rows && rows.length > 0 ? (
            <TableBody>{body}</TableBody>
          ) : (
            <TableBody>
              <TableRow>
                <TableCell colSpan={header.length}>No results found.</TableCell>
              </TableRow>
            </TableBody>
          )}
        </Table>
      </TableContainer>
      <Box
        sx={{
          display: 'flex',
          mt: 2,
          justifyContent: 'justify-between',
          alignItems: 'center',
        }}
      >
        <Button
          sx={{ display: 'flex' }}
          startIcon={<ArrowBackIcon />}
          disabled={!showPrevious}
          onClick={handlePreviousPage}
        >
          Previous page
        </Button>
        <Box sx={{ flexGrow: 1, textAlign: 'center' }}>
          {loading ? (
            <span>Loading...</span>
          ) : (
            <span>Page {pageTokenHistory.length + 1}</span>
          )}
        </Box>
        <Box sx={{ display: 'flex', ml: 'auto' }}>
          <FormControl
            sx={{ m: 1, minWidth: 120 }}
            variant="outlined"
            size="small"
          >
            <InputLabel id="page-size-label">Page size</InputLabel>
            <Select
              id="page-size"
              labelId="page-size-label"
              label="Page size"
              value={rowsPerPage.toString()}
              onChange={(event: SelectChangeEvent) =>
                setRowsPerPage(parseInt(event.target.value))
              }
            >
              {allowedPageSizes.map((pageSize) => (
                <MenuItem key={pageSize.toString()} value={pageSize.toString()}>
                  {pageSize}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <Button
            endIcon={<ArrowForwardIcon />}
            disabled={!showNext}
            onClick={handleNextPage}
          >
            Next page
          </Button>
        </Box>
      </Box>
    </>
  );
}
