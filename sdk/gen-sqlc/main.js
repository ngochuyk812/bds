const fs = require('fs');

// Sinh câu lệnh INSERT
const generateInsertQuery = (table, columns) => {
  const insertColumns = columns.filter(col => col !== "deletedAt" && col !== "updatedAt"); // Giữ lại createdAt
  const columnNames = insertColumns.join(", ");
  const valuePlaceholders = insertColumns.map(() => "?").join(", ");
  return `-- name: Create${capitalize(table)} :exec\nINSERT INTO ${table}s (${columnNames}) VALUES (${valuePlaceholders});`;
};

// Sinh câu lệnh SELECT
const generateSelectQueryById = (table) =>
  `-- name: Get${capitalize(table)}ById :one\nSELECT * FROM ${table}s WHERE id = ?;`;

const generateSelectQueryByGuid = (table) =>
  `-- name: Get${capitalize(table)}ByGuid :one\nSELECT * FROM ${table}s WHERE guid = ?;`;

// Sinh câu lệnh UPDATE by ID
const generateUpdateQuery = (table, columns) => {
  const editableCols = columns.filter(col => !["id", "guid", "createdAt", "deletedAt"].includes(col));
  const setStatements = editableCols.map((col, idx) => `${col} = ?`);
  return `-- name: Update${capitalize(table)}ById :exec\nUPDATE ${table}s SET ${setStatements.join(", ")} WHERE id = ?;`;
};

// Sinh câu lệnh UPDATE by Guid
const generateUpdateByGuidQuery = (table, columns) => {
  const editableCols = columns.filter(col => !["id", "guid", "createdAt", "deletedAt"].includes(col));
  const setStatements = editableCols.map((col, idx) => `${col} = ?`);
  return `-- name: Update${capitalize(table)}ByGuid :exec\nUPDATE ${table}s SET ${setStatements.join(", ")} WHERE guid = ?;`;
};

// Sinh câu lệnh DELETE
const generateDeleteQuery = (table) =>
  `-- name: Delete${capitalize(table)}ById :exec\nUPDATE ${table}s SET deletedAt = ? WHERE id = ?;`;

const generateDeleteByGuidQuery = (table) =>
  `-- name: Delete${capitalize(table)}ByGuid :exec\nUPDATE ${table}s SET deletedAt = ? WHERE guid = ?;`;

// Paging query
const generatePagingQuery = (table) =>
  `-- name: Get${capitalize(table)}sPaging :many\nSELECT * FROM ${table}s WHERE deletedAt IS NULL ORDER BY createdAt DESC LIMIT ? OFFSET ?;`;

// Capitalize function
const capitalize = (str) => str.charAt(0).toUpperCase() + str.slice(1);

// Gộp tất cả
const generateAllQueries = (table, columns) => {
  return [
    generateInsertQuery(table, columns),
    generateSelectQueryById(table),
    generateSelectQueryByGuid(table),
    generateUpdateQuery(table, columns),
    generateUpdateByGuidQuery(table, columns),
    generateDeleteQuery(table),
    generateDeleteByGuidQuery(table),
    generatePagingQuery(table)
  ].join("\n\n");
};

// Example
const tableName = "site";
const columns = ["guid", "siteId", "name", "createdAt", "deletedAt", "updatedAt"];

const sqlQueries = generateAllQueries(tableName, columns);

fs.writeFile('queries.sql', sqlQueries, (err) => {
    if (err) {
      console.error('Có lỗi khi ghi file:', err);
    } else {
      console.log('Câu lệnh SQL đã được ghi vào file queries.sql');
    }
  });
