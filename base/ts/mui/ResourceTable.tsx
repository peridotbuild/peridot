/**
 * Copyright 2023 Peridot Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
import TextField from '@mui/material/TextField';
import InputLabel from '@mui/material/InputLabel';

import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import ArrowForwardIcon from '@mui/icons-material/ArrowForward';

import { StandardResource } from '../resource';

export interface ResourceTableField {
  key: string;
  label: string;
}

export interface ResourceTableProps<T> {
  fields: ResourceTableField[];

  // Default filter to start with
  defaultFilter?: string;

  // load is usually the OpenAPI SDK function that loads the resource.
  load(pageSize: number, pageToken?: string, filter?: string): Promise<any>;

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

  let qFilter = search.get('q');
  let initFilter: string | undefined = undefined;
  if (qFilter === null) {
    initFilter = props.defaultFilter;
  } else {
    initFilter = qFilter;
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
  const [filter, setFilter] = React.useState<string | undefined>(initFilter);

  const updateSearch = (replace = false) => {
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

    if (filter) {
      search.set('q', filter);
    } else {
      search.set('q', '');
    }

    // Compare search and only update if different
    // If replace is true, then we should always update
    if (!replace) {
      if ('?' + search.toString() === location.search) {
        return;
      }
    }

    setSearch(search, { replace });
  };

  // First update should replace the query parameters with the initial state
  React.useEffect(() => {
    updateSearch(true);
  }, []);

  // Update state when query parameters change
  // Create a timeout to prevent too many updates
  const searchTimeout = React.useRef<number | undefined>(undefined);
  React.useEffect(() => {
    if (searchTimeout.current) {
      clearTimeout(searchTimeout.current);
    }
    searchTimeout.current = setTimeout(() => {
      const pt = search.get('pt') || undefined;
      const pth = search.get('pth');
      let initPageTokenHistory: string[] = [];
      if (pth) {
        try {
          initPageTokenHistory = JSON.parse(atob(pth));
        } catch (e) {}
      }
      let initRowsPerPage = parseInt(search.get('rpp') || '25') || 25;
      if (!allowedPageSizes.includes(initRowsPerPage)) {
        initRowsPerPage = 25;
      }
      if (pageToken !== pt) {
        setPageToken(pt);
      }
      if (rowsPerPage !== initRowsPerPage) {
        setRowsPerPage(initRowsPerPage);
      }
      if (
        JSON.stringify(pageTokenHistory) !== JSON.stringify(initPageTokenHistory)
      ) {
        setPageTokenHistory(initPageTokenHistory);
      }

      const q = search.get('q') || undefined;
      if (filter !== q) {
        setFilter(q);
      }
    }, 500);
  }, [search]);

  // Update the query parameters when any of the pagination state changes
  React.useEffect(() => {
    updateSearch();
  }, [pageToken, pageTokenHistory, filter]);

  // Rows per page changing means we need to reset the page token history
  React.useEffect(() => {
    const search = new URLSearchParams(location.search);
    search.set('rpp', rowsPerPage.toString());

    setPageTokenHistory([]);
    setPageToken(undefined);

    if ('?' + search.toString() === location.search) {
      return;
    }

    setSearch(search);
  }, [rowsPerPage]);

  const fetchResource = async () => {
    setLoading(true);
    const [res, err] = await props.load(rowsPerPage, pageToken, filter);
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
  };

  // Load the resource using useEffect
  React.useEffect(() => {
    fetchResource().then();
  }, [pageToken, rowsPerPage]);

  // For filter, we're going to wait for the user to stop typing for 500ms
  // before we actually fetch the resource.
  // This is to prevent the resource from being fetched too many times.
  const timeout = React.useRef<number | undefined>(undefined);

  React.useEffect(() => {
    if (timeout.current) {
      clearTimeout(timeout.current);
    }
    timeout.current = setTimeout(() => {
      fetchResource().then();
    }, 500);
  }, [filter]);

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

  // Create a search box for filter input
  // This can be disabled if the request does not support filtering
  const searchBox = (
    <Box sx={{ display: 'flex', alignItems: 'center', mb: 2, width: '100%' }}>
      <TextField
        sx={{ mr: 2, flexGrow: 1 }}
        label="Filter"
        variant="outlined"
        size="small"
        value={filter}
        onChange={(event: React.ChangeEvent<HTMLInputElement>) => setFilter(event.target.value)}
      />
      <Button
        variant="contained"
        onClick={() => {
          setFilter('');
          setPageToken(undefined);
        }}
      >
        Clear
      </Button>
    </Box>
  );

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
      {searchBox}
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
